package service

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/horvatic/zracni-udar-service/pkg/model"
	"github.com/horvatic/zracni-udar-service/pkg/store"
)

type GitRepoService interface {
	GetGitReposByProjectId(id string) ([]model.GitRepo, ErrorType, error)
	CreateGitRepo(projectId string, body *io.ReadCloser) (ErrorType, error)
	UpdateGitRepo(gitRepoId string, projectId string, body *io.ReadCloser) (ErrorType, error)
	DeleteGitRepo(gitRepoId string, projectId string) (ErrorType, error)
	GetGitRepo(projectId string, gitRepoId string) (*model.GitRepo, ErrorType, error)
}

type gitRepoService struct {
	store store.Store
}

func BuildGitRepoService(store store.Store) GitRepoService {
	return &gitRepoService{
		store: store,
	}
}

func (g *gitRepoService) GetGitReposByProjectId(id string) ([]model.GitRepo, ErrorType, error) {
	var gitRepo []model.GitRepo
	project := g.store.GetProjectById(id)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}
	for _, p := range project.ProjectGitRepos {
		gitRepo = append(gitRepo, model.GitRepo{
			Id:          p.Id,
			ProjectId:   project.ProjectId,
			Name:        p.Name,
			Description: p.Description,
			Uri:         p.Uri,
		})
	}
	return gitRepo, NoError, nil
}

func (g *gitRepoService) CreateGitRepo(projectId string, body *io.ReadCloser) (ErrorType, error) {
	var gitRepo model.GitRepo
	err := json.NewDecoder(*body).Decode(&gitRepo)
	if err != nil {
		return JsonError, err
	}
	project := g.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	project.ProjectGitRepos = append(project.ProjectGitRepos, model.ProjectGitRepos{
		Id:          uuid.NewString(),
		Name:        gitRepo.Name,
		Description: gitRepo.Description,
		Uri:         gitRepo.Uri,
	})
	updateErr := g.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (g *gitRepoService) UpdateGitRepo(gitRepoId string, projectId string, body *io.ReadCloser) (ErrorType, error) {
	var gitRepo model.GitRepo
	err := json.NewDecoder(*body).Decode(&gitRepo)
	if err != nil {
		return JsonError, err
	}
	project := g.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectGitRepos {
		if pn.Id == gitRepoId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find git repo")
	}
	project.ProjectGitRepos[index].Name = gitRepo.Name
	project.ProjectGitRepos[index].Description = gitRepo.Description
	project.ProjectGitRepos[index].Uri = gitRepo.Uri
	updateErr := g.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (g *gitRepoService) DeleteGitRepo(gitRepoId string, projectId string) (ErrorType, error) {
	project := g.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectGitRepos {
		if pn.Id == gitRepoId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find git repo")
	}
	project.ProjectGitRepos = append(project.ProjectGitRepos[:index], project.ProjectGitRepos[index+1:]...)

	updateErr := g.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (g *gitRepoService) GetGitRepo(projectId string, gitRepoId string) (*model.GitRepo, ErrorType, error) {
	project := g.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, p := range project.ProjectGitRepos {
		if p.Id == gitRepoId {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, BadRequest, errors.New("can not find git repo")
	}

	return &model.GitRepo{
		Id:          project.ProjectGitRepos[index].Id,
		ProjectId:   project.ProjectId,
		Name:        project.ProjectGitRepos[index].Name,
		Description: project.ProjectGitRepos[index].Description,
		Uri:         project.ProjectGitRepos[index].Uri,
	}, NoError, nil
}
