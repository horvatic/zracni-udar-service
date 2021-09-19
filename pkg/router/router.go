package router

import (
	"github.com/gorilla/mux"

	"github.com/horvatic/zracni-udar-service/pkg/controller"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/projectsmetadata", controller.GetAllProjectsMetaData).Methods("GET")
	r.HandleFunc("/project/{id}", controller.GetProjectById).Methods("GET")
	r.HandleFunc("/project/{id}/notes", controller.GetNotesByProjectId).Methods("GET")
	r.HandleFunc("/project/{id}/blogs", controller.GetBlogsByProjectId).Methods("GET")
	r.HandleFunc("/project/{id}/videos", controller.GetBlogsByProjectId).Methods("GET")
	r.HandleFunc("/project/{id}/diagrams", controller.GetBlogsByProjectId).Methods("GET")
	r.HandleFunc("/project/{id}/gitrepos", controller.GetGitReposByProjectId).Methods("GET")
	r.HandleFunc("/project/{id}/buildmetadatas", controller.GetBuildMetaDatasByProjectId).Methods("GET")
	r.HandleFunc("/project/{projectId}/{buildId}/builds", controller.GetBuildsForProject).Methods("GET")

	return r
}
