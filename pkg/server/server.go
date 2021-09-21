package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/horvatic/zracni-udar-service/pkg/controller"
	"github.com/horvatic/zracni-udar-service/pkg/router"
	"github.com/horvatic/zracni-udar-service/pkg/service"
	"github.com/horvatic/zracni-udar-service/pkg/store"
)

func Start() {
	store := store.BuildMongoProjectStore(os.Getenv("MONGO_CONNECTION_STRING"), os.Getenv("MONGO_DATABASE"), os.Getenv("MONGO_COLLECTION"))
	projectBuildService := service.BuildProjectBuildService()
	projectService := service.BuildProjectService(store)
	projectController := controller.BuildProjectController(projectService, projectBuildService)
	router := router.SetRoutes(projectController)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("FRONT_END_HOST")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
