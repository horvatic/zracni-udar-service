package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horvatic/zracni-udar-service/pkg/service"
)

type ServiceInfoController interface {
	GetServicesInfoByProjectId(w http.ResponseWriter, req *http.Request)
	GetServiceInfo(w http.ResponseWriter, req *http.Request)
	CreateServiceInfo(w http.ResponseWriter, req *http.Request)
	UpdateServiceInfo(w http.ResponseWriter, req *http.Request)
	DeleteServiceInfo(w http.ResponseWriter, req *http.Request)
}

type serviceInfoController struct {
	serviceInfoService service.ServiceInfoService
}

func BuildServiceInfoController(serviceInfoService service.ServiceInfoService) ServiceInfoController {
	return &serviceInfoController{
		serviceInfoService: serviceInfoService,
	}
}

func (b *serviceInfoController) GetServicesInfoByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	servicesInfo, errorType, err := b.serviceInfoService.GetServicesInfoByProjectId(vars["id"])
	sendJson(w, servicesInfo, errorType, err)
}

func (b *serviceInfoController) GetServiceInfo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	serviceInfo, errorType, err := b.serviceInfoService.GetServiceInfo(vars["projectId"], vars["serviceId"])
	sendJson(w, serviceInfo, errorType, err)
}

func (b *serviceInfoController) CreateServiceInfo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := b.serviceInfoService.CreateServiceInfo(vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (b *serviceInfoController) UpdateServiceInfo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := b.serviceInfoService.UpdateServiceInfo(vars["serviceId"], vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (b *serviceInfoController) DeleteServiceInfo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := b.serviceInfoService.DeleteServiceInfo(vars["serviceId"], vars["projectId"])
	sendCreateResult(w, errorType, err)
}
