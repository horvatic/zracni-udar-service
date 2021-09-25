package service

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/horvatic/zracni-udar-service/pkg/model"
	"github.com/horvatic/zracni-udar-service/pkg/store"
)

type ProjectService interface {
	GetProjectsMetaData() ([]model.ProjectMetaData, ErrorType, error)
	GetProjectMetaDataById(id string) (*model.ProjectMetaData, ErrorType, error)
	CreateProject(body *io.ReadCloser) (ErrorType, error)
	UpdateProjectMetaData(id string, body *io.ReadCloser) (ErrorType, error)
	DeleteProject(projectId string) (ErrorType, error)
}

type projectService struct {
	store store.Store
}

func BuildProjectService(store store.Store) ProjectService {
	return &projectService{
		store: store,
	}
}

func (p *projectService) GetProjectsMetaData() ([]model.ProjectMetaData, ErrorType, error) {
	var projectsMetaData []model.ProjectMetaData
	projects := p.store.GetAllProjects()
	for _, p := range projects {
		projectsMetaData = append(projectsMetaData, model.ProjectMetaData{
			Id:          p.ProjectId,
			Name:        p.Name,
			Description: p.Description,
		})
	}
	return projectsMetaData, NoError, nil
}

func (p *projectService) GetProjectMetaDataById(id string) (*model.ProjectMetaData, ErrorType, error) {
	project := p.store.GetProjectById(id)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}
	return &model.ProjectMetaData{
		Id:          project.ProjectId,
		Name:        project.Name,
		Description: project.Description,
	}, NoError, nil
}

func (p *projectService) CreateProject(body *io.ReadCloser) (ErrorType, error) {
	var projectMetaData model.ProjectMetaData
	err := json.NewDecoder(*body).Decode(&projectMetaData)
	if err != nil {
		return JsonError, err
	}
	project := &model.Project{
		ProjectId:   uuid.NewString(),
		Name:        projectMetaData.Name,
		Description: projectMetaData.Description,
	}
	dbErr := p.store.CreateProject(project)
	if dbErr != nil {
		return DatabaseError, dbErr
	}
	return NoError, nil
}

func (p *projectService) UpdateProjectMetaData(id string, body *io.ReadCloser) (ErrorType, error) {
	var projectMetaData model.ProjectMetaData
	err := json.NewDecoder(*body).Decode(&projectMetaData)
	if err != nil {
		return JsonError, err
	}
	project := p.store.GetProjectById(id)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	project.Name = projectMetaData.Name
	project.Description = projectMetaData.Description
	updateErr := p.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (p *projectService) DeleteProject(projectId string) (ErrorType, error) {
	err := p.store.DeleteProject(projectId)
	if err != nil {
		return DatabaseError, err
	}
	return NoError, nil
}
