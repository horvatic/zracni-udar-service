package service

import (
	"encoding/json"
	"errors"
	"io"

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
	GetBuildsForProject(projectid string, buildId string) []model.Build
}

type projectBuildService struct {
	store store.Store
}

func BuildProjectBuildService(store store.Store) ProjectBuildService {
	return &projectBuildService{
		store: store,
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
	}, NoError, nil
}

func (b *projectBuildService) GetBuildsForProject(projectid string, buildId string) []model.Build {
	if projectid == "785588cf-9ec0-4482-9a5f-df343cec6ac4" && buildId == "1b6086ae-a5b8-499f-86ce-e59fc3f84194" {
		return []model.Build{
			{
				Id:        "cf22290b-c20e-4994-abce-112795e32886",
				ProjectId: "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Version:   "0",
				Environments: []model.Environment{
					{
						Name:   "dev",
						Status: "passing",
						Order:  0,
					},
					{
						Name:   "test",
						Status: "passing",
						Order:  0,
					},
					{
						Name:   "prod",
						Status: "passing",
						Order:  0,
					},
				},
			},
			{
				Id:        "16d01ace-27ac-4aa0-b3f3-b3a62a5a7a81",
				ProjectId: "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Version:   "1",
				Environments: []model.Environment{
					{
						Name:   "dev",
						Status: "passing",
						Order:  0,
					},
					{
						Name:   "test",
						Status: "pending",
						Order:  0,
					},
					{
						Name:   "prod",
						Status: "pending",
						Order:  0,
					},
				},
			},
		}
	} else if projectid == "785588cf-9ec0-4482-9a5f-df343cec6ac4" && buildId == "217c4d79-e5a3-4bf2-8beb-4fbd8f556c7e" {
		return []model.Build{
			{
				Id:        "da663d33-9df1-4546-aafa-96f79aa75137",
				ProjectId: "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Version:   "0",
				Environments: []model.Environment{
					{
						Name:   "dev",
						Status: "passing",
						Order:  0,
					},
					{
						Name:   "test",
						Status: "pending",
						Order:  0,
					},
					{
						Name:   "prod",
						Status: "pending",
						Order:  0,
					},
				},
			},
		}
	} else if projectid == "8c98d8dd-28a0-4d2e-896b-31b1a59fed93" && buildId == "3649e08e-2b55-4959-a031-fb26b43f281a" {
		return []model.Build{
			{
				Id:        "3952c302-8dd0-4b7b-9741-d325cd05dfde",
				ProjectId: "8c98d8dd-28a0-4d2e-896b-31b1a59fed93",
				Version:   "0",
				Environments: []model.Environment{
					{
						Name:   "dev",
						Status: "passing",
						Order:  0,
					},
					{
						Name:   "test",
						Status: "failed",
						Order:  0,
					},
					{
						Name:   "prod",
						Status: "pending",
						Order:  0,
					},
				},
			},
		}
	}
	return []model.Build{}
}
