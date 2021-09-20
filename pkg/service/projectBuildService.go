package service

import (
	"github.com/horvatic/zracni-udar-service/pkg/model"
)

type ProjectBuildService interface {
	GetBuildsForProject(projectid string, buildId string) []model.Build
}

type projectBuildService struct {
}

func BuildProjectBuildService() ProjectBuildService {
	return &projectBuildService{}
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
