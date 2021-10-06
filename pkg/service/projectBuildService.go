package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/google/go-github/v39/github"
	"github.com/google/uuid"
	"github.com/horvatic/zracni-udar-service/pkg/model"
	"github.com/horvatic/zracni-udar-service/pkg/store"
)

type ProjectBuildService interface {
	GetBuildMetaDatasByProjectId(id string) ([]model.BuildMetaData, ErrorType, error)
	CreateBuildMetaData(projectId string, body *io.ReadCloser) (ErrorType, error)
	UpdateBuildMetaData(buildId string, projectId string, body *io.ReadCloser) (ErrorType, error)
	DeleteBuildMetaData(buildId string, projectId string) (ErrorType, error)
	GetBuildMetaData(projectId string, buildId string) (*model.BuildMetaData, ErrorType, error)
	GetBuildsForProject(projectid string, buildId string) ([]model.Build, ErrorType, error)
}

type projectBuildService struct {
	store  store.Store
	client *github.Client
}

func BuildProjectBuildService(store store.Store, client *github.Client) ProjectBuildService {
	return &projectBuildService{
		store:  store,
		client: client,
	}
}

func (b *projectBuildService) GetBuildMetaDatasByProjectId(id string) ([]model.BuildMetaData, ErrorType, error) {
	var buildMetaData []model.BuildMetaData
	project := b.store.GetProjectById(id)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}
	for _, p := range project.ProjectBuildsMetaData {
		buildMetaData = append(buildMetaData, model.BuildMetaData{
			Id:          p.Id,
			ProjectId:   project.ProjectId,
			Name:        p.Name,
			Description: p.Description,
			Uri:         p.Uri,
			RepoName:    p.RepoName,
			RepoOwner:   p.RepoOwner,
		})
	}
	return buildMetaData, NoError, nil
}

func (b *projectBuildService) CreateBuildMetaData(projectId string, body *io.ReadCloser) (ErrorType, error) {
	var build model.BuildMetaData
	err := json.NewDecoder(*body).Decode(&build)
	if err != nil {
		return JsonError, err
	}
	project := b.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	project.ProjectBuildsMetaData = append(project.ProjectBuildsMetaData, model.ProjectBuildsMetaData{
		Id:          uuid.NewString(),
		Name:        build.Name,
		Description: build.Description,
		Uri:         build.Uri,
		RepoName:    build.RepoName,
		RepoOwner:   build.RepoOwner,
	})
	updateErr := b.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (b *projectBuildService) UpdateBuildMetaData(buildId string, projectId string, body *io.ReadCloser) (ErrorType, error) {
	var build model.BuildMetaData
	err := json.NewDecoder(*body).Decode(&build)
	if err != nil {
		return JsonError, err
	}
	project := b.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectBuildsMetaData {
		if pn.Id == buildId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find build")
	}
	project.ProjectBuildsMetaData[index].Name = build.Name
	project.ProjectBuildsMetaData[index].Description = build.Description
	project.ProjectBuildsMetaData[index].Uri = build.Uri
	project.ProjectBuildsMetaData[index].RepoName = build.RepoName
	project.ProjectBuildsMetaData[index].RepoOwner = build.RepoOwner
	updateErr := b.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (b *projectBuildService) DeleteBuildMetaData(buildId string, projectId string) (ErrorType, error) {
	project := b.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectBuildsMetaData {
		if pn.Id == buildId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find build")
	}
	project.ProjectBuildsMetaData = append(project.ProjectBuildsMetaData[:index], project.ProjectBuildsMetaData[index+1:]...)

	updateErr := b.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (b *projectBuildService) GetBuildMetaData(projectId string, buildId string) (*model.BuildMetaData, ErrorType, error) {
	project := b.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, p := range project.ProjectBuildsMetaData {
		if p.Id == buildId {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, BadRequest, errors.New("can not find build")
	}

	return &model.BuildMetaData{
		Id:          project.ProjectBuildsMetaData[index].Id,
		ProjectId:   project.ProjectId,
		Name:        project.ProjectBuildsMetaData[index].Name,
		Description: project.ProjectBuildsMetaData[index].Description,
		Uri:         project.ProjectBuildsMetaData[index].Uri,
		RepoName:    project.ProjectBuildsMetaData[index].RepoName,
		RepoOwner:   project.ProjectBuildsMetaData[index].RepoOwner,
	}, NoError, nil
}

func (b *projectBuildService) GetBuildsForProject(projectId string, buildId string) ([]model.Build, ErrorType, error) {
	project := b.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return []model.Build{}, BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, p := range project.ProjectBuildsMetaData {
		if p.Id == buildId {
			index = i
			break
		}
	}

	if index == -1 {
		return []model.Build{}, BadRequest, errors.New("can not find build")
	}

	buildMetaData := project.ProjectBuildsMetaData[index]
	var builds []model.Build
	if workflows, _, err := b.client.Actions.ListWorkflows(context.TODO(), buildMetaData.RepoOwner, buildMetaData.RepoName, &github.ListOptions{}); err != nil {
		return []model.Build{}, BadRequest, err
	} else {
		for _, w := range workflows.Workflows {
			if runs, _, err := b.client.Actions.ListWorkflowRunsByID(context.TODO(), buildMetaData.RepoOwner, buildMetaData.RepoName, w.GetID(), &github.ListWorkflowRunsOptions{}); err != nil {
				return []model.Build{}, BadRequest, err
			} else {
				numRuns := len(runs.WorkflowRuns)
				for i, r := range runs.WorkflowRuns {
					build := model.Build{
						Id:        buildMetaData.Id,
						ProjectId: projectId,
						BuildId:   buildId,
						Version:   strconv.Itoa(numRuns - i),
					}
					if jobs, _, err := b.client.Actions.ListWorkflowJobs(context.TODO(), buildMetaData.RepoOwner, buildMetaData.RepoName, r.GetID(), &github.ListWorkflowJobsOptions{}); err != nil {
						return []model.Build{}, BadRequest, err
					} else {
						for ij, j := range jobs.Jobs {
							build.Stages = append(build.Stages, model.Stage{
								Name:   j.GetName(),
								Status: j.GetStatus(),
								Order:  ij,
							})
						}
					}
					builds = append(builds, build)
				}
			}
		}
	}

	return builds, NoError, nil
}
