package router

import (
	"github.com/gorilla/mux"

	"github.com/horvatic/zracni-udar-service/pkg/controller"
)

func SetRoutes(c controller.ProjectController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/projectsmetadata", c.GetAllProjectsMetaData).Methods("GET")
	r.HandleFunc("/project/{id}", c.GetProjectById).Methods("GET")
	r.HandleFunc("/project/{id}/notes", c.GetNotesByProjectId).Methods("GET")
	r.HandleFunc("/project/{id}/blogs", c.GetBlogsByProjectId).Methods("GET")
	r.HandleFunc("/project/{id}/videos", c.GetVideosByProjectId).Methods("GET")
	r.HandleFunc("/project/{id}/diagrams", c.GetDiagramsByProjectId).Methods("GET")
	r.HandleFunc("/project/{id}/gitrepos", c.GetGitReposByProjectId).Methods("GET")
	r.HandleFunc("/project/{id}/buildmetadatas", c.GetBuildMetaDatasByProjectId).Methods("GET")
	r.HandleFunc("/project/{projectId}/{buildId}/builds", c.GetBuildsForProject).Methods("GET")

	return r
}
