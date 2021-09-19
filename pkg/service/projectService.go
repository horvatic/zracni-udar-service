package service

import (
	"github.com/horvatic/zracni-udar-service/pkg/model"
	"github.com/horvatic/zracni-udar-service/pkg/store"
)

func GetProjectsMetaData() []model.ProjectMetaData {
	return store.GetProjectsMetaData()
}

func GetProjectById(id string) model.ProjectMetaData {
	return store.GetProjectById(id)
}

func GetNotesByProjectId(id string) []model.Note {
	return store.GetNotesByProjectId(id)
}

func GetBlogsByProjectId(id string) []model.Blog {
	return store.GetBlogsByProjectId(id)
}

func GetVideosByProjectId(id string) []model.Video {
	return store.GetVideosByProjectId(id)
}

func GetDiagramsByProjectId(id string) []model.Diagram {
	return store.GetDiagramsByProjectId(id)
}

func GetGitReposByProjectId(id string) []model.GitRepo {
	return store.GetGitReposByProjectId(id)
}

func GetBuildMetaDatasByProjectId(id string) []model.BuildMetaData {
	return store.GetBuildMetaDatasByProjectId(id)
}

func GetBuildsForProject(projectId string, buildId string) []model.Build {
	return store.GetBuildsForProject(projectId, buildId)
}
