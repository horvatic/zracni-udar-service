package service

import (
	"github.com/horvatic/zracni-udar-service/pkg/model"
	"github.com/horvatic/zracni-udar-service/pkg/store"
)

type ProjectService interface {
	GetProjectsMetaData() []model.ProjectMetaData
	GetProjectById(id string) model.ProjectMetaData
	GetNotesByProjectId(id string) []model.Note
	GetBlogsByProjectId(id string) []model.Blog
	GetVideosByProjectId(id string) []model.Video
	GetDiagramsByProjectId(id string) []model.Diagram
	GetGitReposByProjectId(id string) []model.GitRepo
	GetBuildMetaDatasByProjectId(id string) []model.BuildMetaData
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

func (p *projectService) GetProjectById(id string) model.ProjectMetaData {
	return p.store.GetProjectById(id)
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
