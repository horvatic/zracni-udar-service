package service

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/horvatic/zracni-udar-service/pkg/model"
	"github.com/horvatic/zracni-udar-service/pkg/store"
)

type ServiceInfoService interface {
	GetServicesInfoByProjectId(id string) ([]model.ServiceInfo, ErrorType, error)
	CreateServiceInfo(projectId string, body *io.ReadCloser) (ErrorType, error)
	UpdateServiceInfo(serviceInfoId string, projectId string, body *io.ReadCloser) (ErrorType, error)
	DeleteServiceInfo(serviceInfoId string, projectId string) (ErrorType, error)
	GetServiceInfo(projectId string, serviceInfoId string) (*model.ServiceInfo, ErrorType, error)
}

type serviceInfoService struct {
	store store.Store
}

func BuildServiceInfoService(store store.Store) ServiceInfoService {
	return &serviceInfoService{
		store: store,
	}
}

func (b *serviceInfoService) GetServicesInfoByProjectId(id string) ([]model.ServiceInfo, ErrorType, error) {
	var serviceInfos []model.ServiceInfo
	project := b.store.GetProjectById(id)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}

	for _, p := range project.ProjectServicesInfo {
		serviceInfos = append(serviceInfos, model.ServiceInfo{
			Id:          p.Id,
			ProjectId:   project.ProjectId,
			Name:        p.Name,
			Description: p.Description,
			Uri:         p.Uri,
			HealthUri:   p.HealthUri,
		})
	}
	return serviceInfos, NoError, nil
}

func (b *serviceInfoService) CreateServiceInfo(projectId string, body *io.ReadCloser) (ErrorType, error) {
	var serviceInfo model.ServiceInfo
	err := json.NewDecoder(*body).Decode(&serviceInfo)
	if err != nil {
		return JsonError, err
	}
	project := b.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	project.ProjectServicesInfo = append(project.ProjectServicesInfo, model.ProjectServicesInfo{
		Id:          uuid.NewString(),
		Name:        serviceInfo.Name,
		Description: serviceInfo.Description,
		Uri:         serviceInfo.Uri,
		HealthUri:   serviceInfo.HealthUri,
	})
	updateErr := b.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (b *serviceInfoService) UpdateServiceInfo(serviceInfoId string, projectId string, body *io.ReadCloser) (ErrorType, error) {
	var serviceInfo model.ServiceInfo
	err := json.NewDecoder(*body).Decode(&serviceInfo)
	if err != nil {
		return JsonError, err
	}
	project := b.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectServicesInfo {
		if pn.Id == serviceInfoId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find serviceInfo")
	}
	project.ProjectServicesInfo[index].Name = serviceInfo.Name
	project.ProjectServicesInfo[index].Description = serviceInfo.Description
	project.ProjectServicesInfo[index].Uri = serviceInfo.Uri
	project.ProjectServicesInfo[index].HealthUri = serviceInfo.HealthUri
	updateErr := b.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (b *serviceInfoService) DeleteServiceInfo(serviceInfoId string, projectId string) (ErrorType, error) {
	project := b.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectServicesInfo {
		if pn.Id == serviceInfoId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find serviceInfo")
	}
	project.ProjectServicesInfo = append(project.ProjectServicesInfo[:index], project.ProjectServicesInfo[index+1:]...)

	updateErr := b.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (b *serviceInfoService) GetServiceInfo(projectId string, serviceInfoId string) (*model.ServiceInfo, ErrorType, error) {
	project := b.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, p := range project.ProjectServicesInfo {
		if p.Id == serviceInfoId {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, BadRequest, errors.New("can not find serviceInfo")
	}

	return &model.ServiceInfo{
		Id:          project.ProjectServicesInfo[index].Id,
		ProjectId:   project.ProjectId,
		Name:        project.ProjectServicesInfo[index].Name,
		Description: project.ProjectServicesInfo[index].Description,
		Uri:         project.ProjectServicesInfo[index].Uri,
		HealthUri:   project.ProjectServicesInfo[index].HealthUri,
	}, NoError, nil
}
