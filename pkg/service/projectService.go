package service

import (
	"encoding/json"
	"io"

	"github.com/horvatic/zracni-udar-service/pkg/model"
	"github.com/horvatic/zracni-udar-service/pkg/store"
)

type ErrorType int64

const (
	DatabaseError ErrorType = iota
	JsonError
	BadRequest
	NoError
)

type ProjectService interface {
	GetProjectsMetaData() []model.ProjectMetaData
	GetProjectMetaDataById(id string) *model.ProjectMetaData
	GetNotesByProjectId(id string) []model.Note
	GetBlogsByProjectId(id string) []model.Blog
	GetVideosByProjectId(id string) []model.Video
	GetDiagramsByProjectId(id string) []model.Diagram
	GetGitReposByProjectId(id string) []model.GitRepo
	GetBuildMetaDatasByProjectId(id string) []model.BuildMetaData
	CreateProject(body *io.ReadCloser) (ErrorType, error)
	UpdateProjectMetaData(id string, body *io.ReadCloser) (ErrorType, error)
}

type projectService struct {
	store store.Store
}

func BuildProjectService(store store.Store) ProjectService {
	return &projectService{
		store: store,
	}
}

func (p *projectService) GetProjectsMetaData() []model.ProjectMetaData {
	return p.store.GetProjectsMetaData()
}

func (p *projectService) GetProjectMetaDataById(id string) *model.ProjectMetaData {
	return p.store.GetProjectMetaDataById(id)
}

func (p *projectService) GetNotesByProjectId(id string) []model.Note {
	return p.store.GetNotesByProjectId(id)
}

func (p *projectService) GetBlogsByProjectId(id string) []model.Blog {
	return p.store.GetBlogsByProjectId(id)
}

func (p *projectService) GetVideosByProjectId(id string) []model.Video {
	return p.store.GetVideosByProjectId(id)
}

func (p *projectService) GetDiagramsByProjectId(id string) []model.Diagram {
	return p.store.GetDiagramsByProjectId(id)
}

func (p *projectService) GetGitReposByProjectId(id string) []model.GitRepo {
	return p.store.GetGitReposByProjectId(id)
}

func (p *projectService) GetBuildMetaDatasByProjectId(id string) []model.BuildMetaData {
	return p.store.GetBuildMetaDatasByProjectId(id)
}

func (p *projectService) CreateProject(body *io.ReadCloser) (ErrorType, error) {
	var projectMetaData model.ProjectMetaData
	err := json.NewDecoder(*body).Decode(&projectMetaData)
	if err != nil {
		return JsonError, err
	}
	dbErr := p.store.CreateProject(&projectMetaData)
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
		return BadRequest, err
	}
	project.Name = projectMetaData.Name
	project.Description = projectMetaData.Description
	updateErr := p.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}
