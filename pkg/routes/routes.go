package routes

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/horvatic/zracni-udar-service/pkg/controller"
)

func SetProjectRoutes(c controller.ProjectController, router *mux.Router, namespace string, serviceName string) *mux.Router {
	router.HandleFunc(buildRoute(namespace, serviceName, "/projectsmetadata"), c.GetAllProjectsMetaData).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/projectmetadata"), c.CreateProject).Methods("POST")
	router.HandleFunc(buildRoute(namespace, serviceName, "/projectmetadata/{id}"), c.UpdateProject).Methods("PATCH")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{id}"), c.GetProjectMetaDataById).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{id}"), c.DeleteProject).Methods("DELETE")
	return router
}

func SetNoteRoutes(c controller.NoteController, router *mux.Router, namespace string, serviceName string) *mux.Router {
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{id}/notes"), c.GetNotesByProjectId).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/notes/{noteId}"), c.GetNote).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/notes"), c.CreateNote).Methods("POST")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/notes/{noteId}"), c.UpdateNote).Methods("PATCH")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/notes/{noteId}"), c.DeleteNote).Methods("DELETE")
	return router
}

func SetBlogRoutes(c controller.BlogController, router *mux.Router, namespace string, serviceName string) *mux.Router {
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{id}/blogs"), c.GetBlogsByProjectId).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/blogs/{blogId}"), c.GetBlog).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/blogs"), c.CreateBlog).Methods("POST")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/blogs/{blogId}"), c.UpdateBlog).Methods("PATCH")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/blogs/{blogId}"), c.DeleteBlog).Methods("DELETE")
	return router
}

func SetVideoRoutes(c controller.VideoController, router *mux.Router, namespace string, serviceName string) *mux.Router {
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{id}/videos"), c.GetVideosByProjectId).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/videos/{videoId}"), c.GetVideo).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/videos"), c.CreateVideo).Methods("POST")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/videos/{videoId}"), c.UpdateVideo).Methods("PATCH")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/videos/{videoId}"), c.DeleteVideo).Methods("DELETE")
	return router
}

func SetDiagramRoutes(c controller.DiagramController, router *mux.Router, namespace string, serviceName string) *mux.Router {
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{id}/diagrams"), c.GetDiagramsByProjectId).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/diagrams/{diagramId}"), c.GetDiagram).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/diagrams"), c.CreateDiagram).Methods("POST")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/diagrams/{diagramId}"), c.UpdateDiagram).Methods("PATCH")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/diagrams/{diagramId}"), c.DeleteDiagram).Methods("DELETE")
	return router
}

func SetGitRepoRoutes(c controller.GitRepoController, router *mux.Router, namespace string, serviceName string) *mux.Router {
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{id}/gitrepos"), c.GetGitReposByProjectId).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/gitrepos/{gitRepoId}"), c.GetGitRepo).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/gitrepos"), c.CreateGitRepo).Methods("POST")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/gitrepos/{gitRepoId}"), c.UpdateGitRepo).Methods("PATCH")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/gitrepos/{gitRepoId}"), c.DeleteGitRepo).Methods("DELETE")
	return router
}

func SetProjectBuildRoutes(c controller.ProjectBuildController, router *mux.Router, namespace string, serviceName string) *mux.Router {
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{id}/buildmetadatas"), c.GetBuildMetaDatasByProjectId).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/buildmetadatas/{buildId}"), c.GetBuildMetaData).Methods("GET")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/buildmetadatas"), c.CreateBuildMetaData).Methods("POST")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/buildmetadatas/{buildId}"), c.UpdateBuildMetaData).Methods("PATCH")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/buildmetadatas/{buildId}"), c.DeleteBuildMetaData).Methods("DELETE")
	router.HandleFunc(buildRoute(namespace, serviceName, "/project/{projectId}/{buildId}/builds"), c.GetBuildsForProject).Methods("GET")
	return router
}

func SetHeathRoutes(router *mux.Router, namespace string, serviceName string) *mux.Router {
	router.HandleFunc(buildRoute(namespace, serviceName, "/health"), controller.Health).Methods("GET")
	return router
}

func buildRoute(namespace string, serviceName string, baseRoute string) string {
	if namespace == "" || serviceName == "" {
		return baseRoute
	}
	return fmt.Sprintf("/%s/%s%s", namespace, serviceName, baseRoute)
}
