package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/go-github/v39/github"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/horvatic/zracni-udar-service/pkg/controller"
	"github.com/horvatic/zracni-udar-service/pkg/routes"
	"github.com/horvatic/zracni-udar-service/pkg/service"
	"github.com/horvatic/zracni-udar-service/pkg/store"
	"golang.org/x/oauth2"
)

func Start() {
	router := mux.NewRouter()
	store, dbClient, dbContext, err := store.BuildMongoProjectStore(os.Getenv("MONGO_CONNECTION_STRING"), os.Getenv("MONGO_DATABASE"), os.Getenv("MONGO_COLLECTION"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_PAT")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	routes.SetProjectRoutes(controller.BuildProjectController(service.BuildProjectService(store)), router)
	routes.SetNoteRoutes(controller.BuildNoteController(service.BuildNoteService(store)), router)
	routes.SetBlogRoutes(controller.BuildBlogController(service.BuildBlogService(store)), router)
	routes.SetDiagramRoutes(controller.BuildDiagramController(service.BuildDiagramService(store)), router)
	routes.SetGitRepoRoutes(controller.BuildGitRepoController(service.BuildGitRepoService(store)), router)
	routes.SetProjectBuildRoutes(controller.BuildProjectBuildController(service.BuildProjectBuildService(store, client)), router)
	routes.SetVideoRoutes(controller.BuildVideoController(service.BuildVideoService(store)), router)
	routes.SetServiceRoutes(controller.BuildServiceInfoController(service.BuildServiceInfoService(store)), router)
	routes.SetHeathRoutes(router)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("FRONT_END_HOST")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"})

	fmt.Println("Starting Server On Port 8080")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{Addr: ":8080", Handler: handlers.CORS(originsOk, headersOk, methodsOk)(router)}
	go func() {
		server.ListenAndServe()
	}()

	<-done
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Server Stopped")
	dbClient.Disconnect(dbContext)
	fmt.Println("Db Disconnected")
}
