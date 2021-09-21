package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horvatic/zracni-udar-service/pkg/service"
)

type ProjectController interface {
	GetAllProjectsMetaData(w http.ResponseWriter, req *http.Request)
	GetProjectMetaDataById(w http.ResponseWriter, req *http.Request)
	GetNotesByProjectId(w http.ResponseWriter, req *http.Request)
	GetBlogsByProjectId(w http.ResponseWriter, req *http.Request)
	GetVideosByProjectId(w http.ResponseWriter, req *http.Request)
	GetDiagramsByProjectId(w http.ResponseWriter, req *http.Request)
	GetGitReposByProjectId(w http.ResponseWriter, req *http.Request)
	GetBuildMetaDatasByProjectId(w http.ResponseWriter, req *http.Request)
	GetBuildsForProject(w http.ResponseWriter, req *http.Request)
	CreateProject(w http.ResponseWriter, req *http.Request)
	UpdateProject(w http.ResponseWriter, req *http.Request)
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

func (pc *projectController) GetProjectMetaDataById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	projects := pc.projectService.GetProjectMetaDataById(vars["id"])
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

func (pc *projectController) CreateProject(w http.ResponseWriter, req *http.Request) {
	errorType, err := pc.projectService.CreateProject(&req.Body)
	sendCreateResult(w, errorType, err)
}

func (pc *projectController) UpdateProject(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := pc.projectService.UpdateProjectMetaData(vars["id"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func sendCreateResult(w http.ResponseWriter, errorType service.ErrorType, err error) {
	if errorType == service.JsonError || errorType == service.BadRequest {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if errorType == service.DatabaseError {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func sendJson(w http.ResponseWriter, j interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(j)
}
