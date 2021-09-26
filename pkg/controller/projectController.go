package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horvatic/zracni-udar-service/pkg/service"
)

type ProjectController interface {
	GetAllProjectsMetaData(w http.ResponseWriter, req *http.Request)
	GetProjectMetaDataById(w http.ResponseWriter, req *http.Request)
	CreateProject(w http.ResponseWriter, req *http.Request)
	UpdateProject(w http.ResponseWriter, req *http.Request)
	DeleteProject(w http.ResponseWriter, req *http.Request)
}

type projectController struct {
	projectService service.ProjectService
}

func BuildProjectController(projectService service.ProjectService) ProjectController {
	return &projectController{
		projectService: projectService,
	}
}

func (pc *projectController) GetAllProjectsMetaData(w http.ResponseWriter, req *http.Request) {
	projectMetaData, errorType, err := pc.projectService.GetProjectsMetaData()
	sendJson(w, projectMetaData, errorType, err)
}

func (pc *projectController) GetProjectMetaDataById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	projects, errorType, err := pc.projectService.GetProjectMetaDataById(vars["id"])
	sendJson(w, projects, errorType, err)
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

func (pc *projectController) DeleteProject(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := pc.projectService.DeleteProject(vars["id"])
	sendCreateResult(w, errorType, err)
}
