package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horvatic/zracni-udar-service/pkg/service"
)

type DiagramController interface {
	GetDiagramsByProjectId(w http.ResponseWriter, req *http.Request)
	GetDiagram(w http.ResponseWriter, req *http.Request)
	CreateDiagram(w http.ResponseWriter, req *http.Request)
	UpdateDiagram(w http.ResponseWriter, req *http.Request)
	DeleteDiagram(w http.ResponseWriter, req *http.Request)
}

type diagramController struct {
	diagramService service.DiagramService
}

func BuildDiagramController(diagramService service.DiagramService) DiagramController {
	return &diagramController{
		diagramService: diagramService,
	}
}

func (d *diagramController) GetDiagramsByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	diagram, errorType, err := d.diagramService.GetDiagramsByProjectId(vars["id"])
	sendJson(w, diagram, errorType, err)
}

func (d *diagramController) GetDiagram(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	diagram, errorType, err := d.diagramService.GetDiagram(vars["projectId"], vars["diagramId"])
	sendJson(w, diagram, errorType, err)
}

func (d *diagramController) CreateDiagram(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := d.diagramService.CreateDiagram(vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (d *diagramController) UpdateDiagram(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := d.diagramService.UpdateDiagram(vars["diagramId"], vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (d *diagramController) DeleteDiagram(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := d.diagramService.DeleteDiagram(vars["diagramId"], vars["projectId"])
	sendCreateResult(w, errorType, err)
}
