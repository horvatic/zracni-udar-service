package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horvatic/zracni-udar-service/pkg/service"
)

type ProjectController interface {
	GetAllProjectsMetaData(w http.ResponseWriter, req *http.Request)
	GetProjectById(w http.ResponseWriter, req *http.Request)
	GetNotesByProjectId(w http.ResponseWriter, req *http.Request)
	GetBlogsByProjectId(w http.ResponseWriter, req *http.Request)
	GetVideosByProjectId(w http.ResponseWriter, req *http.Request)
	GetDiagramsByProjectId(w http.ResponseWriter, req *http.Request)
	GetGitReposByProjectId(w http.ResponseWriter, req *http.Request)
	GetBuildMetaDatasByProjectId(w http.ResponseWriter, req *http.Request)
	GetBuildsForProject(w http.ResponseWriter, req *http.Request)
}

type projectController struct {
	projectService      service.ProjectService
	projectBuildService service.ProjectBuildService
}

func BuildProjectController(projectService service.ProjectService, projectBuildService service.ProjectBuildService) ProjectController {
	return &projectController{
		projectService:      projectService,
		projectBuildService: projectBuildService,
	}
}

func (pc *projectController) GetAllProjectsMetaData(w http.ResponseWriter, req *http.Request) {
	projectMetaData := pc.projectService.GetProjectsMetaData()
	sendJson(w, projectMetaData)
}

func (pc *projectController) GetProjectById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	projects := pc.projectService.GetProjectById(vars["id"])
	sendJson(w, projects)
}

func (pc *projectController) GetNotesByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	notes := pc.projectService.GetNotesByProjectId(vars["id"])
	sendJson(w, notes)
}

func (pc *projectController) GetBlogsByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	blogs := pc.projectService.GetBlogsByProjectId(vars["id"])
	sendJson(w, blogs)
}

func (pc *projectController) GetVideosByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	videos := pc.projectService.GetVideosByProjectId(vars["id"])
	sendJson(w, videos)
}

func (pc *projectController) GetDiagramsByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	videos := pc.projectService.GetDiagramsByProjectId(vars["id"])
	sendJson(w, videos)
}

func (pc *projectController) GetGitReposByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	gitRepos := pc.projectService.GetGitReposByProjectId(vars["id"])
	sendJson(w, gitRepos)
}

func (pc *projectController) GetBuildMetaDatasByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	builds := pc.projectService.GetBuildMetaDatasByProjectId(vars["id"])
	sendJson(w, builds)
}

func (pc *projectController) GetBuildsForProject(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	deployments := pc.projectBuildService.GetBuildsForProject(vars["projectId"], vars["buildId"])
	sendJson(w, deployments)
}

func sendJson(w http.ResponseWriter, j interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(j)
}
