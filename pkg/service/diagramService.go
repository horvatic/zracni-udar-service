package service

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/horvatic/zracni-udar-service/pkg/model"
	"github.com/horvatic/zracni-udar-service/pkg/store"
)

type DiagramService interface {
	GetDiagramsByProjectId(id string) ([]model.Diagram, ErrorType, error)
	CreateDiagram(projectId string, body *io.ReadCloser) (ErrorType, error)
	UpdateDiagram(diagramId string, projectId string, body *io.ReadCloser) (ErrorType, error)
	DeleteDiagram(diagramId string, projectId string) (ErrorType, error)
	GetDiagram(projectId string, diagramId string) (*model.Diagram, ErrorType, error)
}

type diagramService struct {
	store store.Store
}

func BuildDiagramService(store store.Store) DiagramService {
	return &diagramService{
		store: store,
	}
}

func (d *diagramService) GetDiagramsByProjectId(id string) ([]model.Diagram, ErrorType, error) {
	var diagrams []model.Diagram
	project := d.store.GetProjectById(id)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}
	for _, p := range project.ProjectDiagrams {
		diagrams = append(diagrams, model.Diagram{
			Id:          p.Id,
			ProjectId:   project.ProjectId,
			Name:        p.Name,
			Description: p.Description,
			Uri:         p.Uri,
		})
	}
	return diagrams, NoError, nil
}

func (d *diagramService) CreateDiagram(projectId string, body *io.ReadCloser) (ErrorType, error) {
	var diagram model.Diagram
	err := json.NewDecoder(*body).Decode(&diagram)
	if err != nil {
		return JsonError, err
	}
	project := d.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	project.ProjectDiagrams = append(project.ProjectDiagrams, model.ProjectDiagrams{
		Id:          uuid.NewString(),
		Name:        diagram.Name,
		Description: diagram.Description,
		Uri:         diagram.Uri,
	})
	updateErr := d.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (d *diagramService) UpdateDiagram(diagramId string, projectId string, body *io.ReadCloser) (ErrorType, error) {
	var diagram model.Diagram
	err := json.NewDecoder(*body).Decode(&diagram)
	if err != nil {
		return JsonError, err
	}
	project := d.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectDiagrams {
		if pn.Id == diagramId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find diagram")
	}
	project.ProjectDiagrams[index].Name = diagram.Name
	project.ProjectDiagrams[index].Description = diagram.Description
	project.ProjectDiagrams[index].Uri = diagram.Uri
	updateErr := d.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (d *diagramService) DeleteDiagram(diagramId string, projectId string) (ErrorType, error) {
	project := d.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectDiagrams {
		if pn.Id == diagramId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find diagram")
	}
	project.ProjectDiagrams = append(project.ProjectDiagrams[:index], project.ProjectDiagrams[index+1:]...)

	updateErr := d.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (d *diagramService) GetDiagram(projectId string, diagramId string) (*model.Diagram, ErrorType, error) {
	project := d.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, p := range project.ProjectDiagrams {
		if p.Id == diagramId {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, BadRequest, errors.New("can not find diagram")
	}

	return &model.Diagram{
		Id:          project.ProjectDiagrams[index].Id,
		ProjectId:   project.ProjectId,
		Name:        project.ProjectDiagrams[index].Name,
		Description: project.ProjectDiagrams[index].Description,
		Uri:         project.ProjectDiagrams[index].Uri,
	}, NoError, nil
}
