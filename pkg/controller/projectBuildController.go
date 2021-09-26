package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horvatic/zracni-udar-service/pkg/service"
)

type ProjectBuildController interface {
	GetBuildsForProject(w http.ResponseWriter, req *http.Request)
	CreateBuildMetaData(w http.ResponseWriter, req *http.Request)
	UpdateBuildMetaData(w http.ResponseWriter, req *http.Request)
	DeleteBuildMetaData(w http.ResponseWriter, req *http.Request)
	GetBuildMetaData(w http.ResponseWriter, req *http.Request)
	GetBuildMetaDatasByProjectId(w http.ResponseWriter, req *http.Request)
}

type projectBuildController struct {
	projectBuildService service.ProjectBuildService
}

func BuildProjectBuildController(projectBuildService service.ProjectBuildService) ProjectBuildController {
	return &projectBuildController{
		projectBuildService: projectBuildService,
	}
}

func (p *projectBuildController) GetBuildsForProject(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	builds := p.projectBuildService.GetBuildsForProject(vars["projectId"], vars["buildId"])
	sendJson(w, builds, service.NoError, nil)
}

func (p *projectBuildController) CreateBuildMetaData(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := p.projectBuildService.CreateBuildMetaData(vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (p *projectBuildController) UpdateBuildMetaData(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := p.projectBuildService.UpdateBuildMetaData(vars["buildId"], vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (p *projectBuildController) DeleteBuildMetaData(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := p.projectBuildService.DeleteBuildMetaData(vars["buildId"], vars["projectId"])
	sendCreateResult(w, errorType, err)
}

func (p *projectBuildController) GetBuildMetaData(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	blog, errorType, err := p.projectBuildService.GetBuildMetaData(vars["projectId"], vars["buildId"])
	sendJson(w, blog, errorType, err)
}

func (p *projectBuildController) GetBuildMetaDatasByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	builds, errorType, err := p.projectBuildService.GetBuildMetaDatasByProjectId(vars["id"])
	sendJson(w, builds, errorType, err)
}
