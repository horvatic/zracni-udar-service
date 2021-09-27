package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/horvatic/zracni-udar-service/pkg/controller"
	"github.com/horvatic/zracni-udar-service/pkg/routes"
	"github.com/horvatic/zracni-udar-service/pkg/service"
	"github.com/horvatic/zracni-udar-service/pkg/store"
)

func Start() {
	namespace := os.Getenv("NAMESPACE")
	serviceName := os.Getenv("SERVICE")
	router := mux.NewRouter()
	store := store.BuildMongoProjectStore(os.Getenv("MONGO_CONNECTION_STRING"), os.Getenv("MONGO_DATABASE"), os.Getenv("MONGO_COLLECTION"))

	routes.SetProjectRoutes(controller.BuildProjectController(service.BuildProjectService(store)), router, namespace, serviceName)
	routes.SetNoteRoutes(controller.BuildNoteController(service.BuildNoteService(store)), router, namespace, serviceName)
	routes.SetBlogRoutes(controller.BuildBlogController(service.BuildBlogService(store)), router, namespace, serviceName)
	routes.SetDiagramRoutes(controller.BuildDiagramController(service.BuildDiagramService(store)), router, namespace, serviceName)
	routes.SetGitRepoRoutes(controller.BuildGitRepoController(service.BuildGitRepoService(store)), router, namespace, serviceName)
	routes.SetProjectBuildRoutes(controller.BuildProjectBuildController(service.BuildProjectBuildService(store)), router, namespace, serviceName)
	routes.SetVideoRoutes(controller.BuildVideoController(service.BuildVideoService(store)), router, namespace, serviceName)
	routes.SetHeathRoutes(router, namespace, serviceName)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("FRONT_END_HOST")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"})

	fmt.Println("Starting Server On Port 8080")
	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
