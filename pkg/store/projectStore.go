package store

import (
	"github.com/horvatic/zracni-udar-service/pkg/model"
)

func GetProjectsMetaData() []model.ProjectMetaData {
	return []model.ProjectMetaData{
		{
			Id:          "785588cf-9ec0-4482-9a5f-df343cec6ac4",
			Name:        "Project 1",
			Description: "Projct description here",
		},
		{
			Id:          "8c98d8dd-28a0-4d2e-896b-31b1a59fed93",
			Name:        "Project 2",
			Description: "description of project 2",
		},
	}
}

func GetProjectById(id string) model.ProjectMetaData {
	if id == "785588cf-9ec0-4482-9a5f-df343cec6ac4" {
		return model.ProjectMetaData{
			Id:          "785588cf-9ec0-4482-9a5f-df343cec6ac4",
			Name:        "Project 1",
			Description: "Projct description here",
		}
	} else if id == "8c98d8dd-28a0-4d2e-896b-31b1a59fed93" {
		return model.ProjectMetaData{
			Id:          "8c98d8dd-28a0-4d2e-896b-31b1a59fed93",
			Name:        "Project 2",
			Description: "description of project 2",
		}
	}
	return model.ProjectMetaData{}
}

func GetNotesByProjectId(id string) []model.Note {
	if id == "785588cf-9ec0-4482-9a5f-df343cec6ac4" {
		return []model.Note{
			{
				Id:        "0071f491-755d-43df-967f-4c60a3b0ca8f",
				ProjectId: "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:      "Project Note 1",
				Note:      "Note:fewfew few ",
			},
			{
				Id:        "4bb2ea6a-0c42-4ced-aae7-0c008f7a559e",
				ProjectId: "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:      "Project Note 2",
				Note:      "Note:fsfsfefe few ",
			},
			{
				Id:        "ee8fc2c2-68c6-47be-adea-a6c5ad7146b6",
				ProjectId: "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:      "Project Note 3",
				Note:      "Note:fewfew ffefew333r3r3 ",
			},
		}
	} else if id == "8c98d8dd-28a0-4d2e-896b-31b1a59fed93" {
		return []model.Note{
			{
				Id:        "8608fc8e-3c8e-48e6-8bd7-1e4a88fb18ff",
				ProjectId: "8c98d8dd-28a0-4d2e-896b-31b1a59fed93",
				Name:      "Project Note 4",
				Note:      "Note: 4 ",
			},
		}
	}
	return []model.Note{}
}

func GetBlogsByProjectId(id string) []model.Blog {
	if id == "785588cf-9ec0-4482-9a5f-df343cec6ac4" {
		return []model.Blog{
			{
				Id:          "b6673aa0-9316-41ef-9d55-8ed646b1eb26",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Blog 1",
				Description: "note info here: fwefweewfewfwefwe",
				Uri:         "https://blog1.horvatic.com",
			},
			{
				Id:          "3017171f-9f84-4005-99c0-95b692898c46",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Blog 2",
				Description: "frrfewf f 34ffeef sf wed23",
				Uri:         "https://blog2.horvatic.com",
			},
			{
				Id:          "41c0dfd5-7c4e-44c1-913f-67c77657e1c8",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Blog 3",
				Description: "rwefhj7uj8ik 8  8 k7j7ujuyj 7uj7",
				Uri:         "https://blog3.horvatic.com",
			},
		}
	} else if id == "8c98d8dd-28a0-4d2e-896b-31b1a59fed93" {
		return []model.Blog{
			{
				Id:          "03bffd1c-5d5a-4d13-bc32-37101085b32a",
				ProjectId:   "8c98d8dd-28a0-4d2e-896b-31b1a59fed93",
				Name:        "Blog 4",
				Description: "tfdg  4t4t fgdrdg4tt h7 h7h7h7h",
				Uri:         "https://blog4.horvatic.com",
			},
		}
	}
	return []model.Blog{}
}

func GetVideosByProjectId(id string) []model.Video {
	if id == "785588cf-9ec0-4482-9a5f-df343cec6ac4" {
		return []model.Video{
			{
				Id:          "cc602e00-e483-4c90-94b5-93b0c342f5bd",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Video 1",
				Description: "note info here: fwefweewfewfwefwe",
				Uri:         "https://www.youtube.com/watch?v=0M8AYU_hPas",
			},
			{
				Id:          "b0faff24-6b41-4bf6-a241-34b23b21c42b",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Video 2",
				Description: "frrfewf f 34ffeef sf wed23",
				Uri:         "https://www.youtube.com/watch?v=tomUWcQ0P3k",
			},
			{
				Id:          "1c2023c3-02cd-424f-bd16-170203cf2500",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Video 3",
				Description: "rwefhj7uj8ik 8  8 k7j7ujuyj 7uj7",
				Uri:         "https://www.youtube.com/watch?v=G1rOthIU-uo",
			},
		}
	} else if id == "8c98d8dd-28a0-4d2e-896b-31b1a59fed93" {
		return []model.Video{
			{
				Id:          "f96be6c2-1011-4dec-bc8b-668ec02ad425",
				ProjectId:   "8c98d8dd-28a0-4d2e-896b-31b1a59fed93",
				Name:        "Video 4",
				Description: "tfdg  4t4t fgdrdg4tt h7 h7h7h7h",
				Uri:         "https://www.youtube.com/watch?v=xiZ61BkMKo8",
			},
		}
	}
	return []model.Video{}
}

func GetDiagramsByProjectId(id string) []model.Diagram {
	if id == "785588cf-9ec0-4482-9a5f-df343cec6ac4" {
		return []model.Diagram{
			{
				Id:          "2138a129-494f-487e-839f-e365ae2ca0a0",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Diagram 1",
				Description: "note info here: fwefweewfewfwefwe",
				Uri:         "https://diagrams1.horvatic.com",
			},
			{
				Id:          "6c83a6a7-1037-4e96-aa6b-08257a9c3c54",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Diagram 2",
				Description: "frrfewf f 34ffeef sf wed23",
				Uri:         "https://diagrams2.horvatic.com",
			},
			{
				Id:          "5293a608-9e11-4eb2-8f22-e67231656dc9",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Diagram 3",
				Description: "rwefhj7uj8ik 8  8 k7j7ujuyj 7uj7",
				Uri:         "https://diagrams3.horvatic.com",
			},
		}
	} else if id == "8c98d8dd-28a0-4d2e-896b-31b1a59fed93" {
		return []model.Diagram{
			{
				Id:          "51d2d98f-c008-4f4b-a911-6fca4ec47891",
				ProjectId:   "8c98d8dd-28a0-4d2e-896b-31b1a59fed93",
				Name:        "Diagram 4",
				Description: "tfdg  4t4t fgdrdg4tt h7 h7h7h7",
				Uri:         "https://diagrams4.horvatic.com",
			},
		}
	}
	return []model.Diagram{}
}

func GetGitReposByProjectId(id string) []model.GitRepo {
	if id == "785588cf-9ec0-4482-9a5f-df343cec6ac4" {
		return []model.GitRepo{
			{
				Id:          "9e78b918-31e0-4398-a8ab-d9d5d1dbd543",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Git Repo 1",
				Description: "note info here: fwefweewfewfwefwe",
				Uri:         "https://github.horvatic.com",
			},
			{
				Id:          "38de8262-6b3f-459c-a9dd-d3ba72ec5aaa",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Git Repo 2",
				Description: "frrfewf f 34ffeef sf wed23",
				Uri:         "https://github2.horvatic.com",
			},
			{
				Id:          "52088195-6c7f-411f-a136-0b3c5308cb27",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Git Repo 3",
				Description: "rwefhj7uj8ik 8  8 k7j7ujuyj 7uj7",
				Uri:         "https://github3.horvatic.com",
			},
		}
	} else if id == "8c98d8dd-28a0-4d2e-896b-31b1a59fed93" {
		return []model.GitRepo{
			{
				Id:          "8227991d-82ea-4cd6-9fd0-3255a12fa951",
				ProjectId:   "8c98d8dd-28a0-4d2e-896b-31b1a59fed93",
				Name:        "Git Repo 4",
				Description: "tfdg  4t4t fgdrdg4tt h7 h7h7h7",
				Uri:         "https://github4.horvatic.com",
			},
		}
	}
	return []model.GitRepo{}
}

func GetBuildMetaDatasByProjectId(id string) []model.BuildMetaData {
	if id == "785588cf-9ec0-4482-9a5f-df343cec6ac4" {
		return []model.BuildMetaData{
			{
				Id:          "1b6086ae-a5b8-499f-86ce-e59fc3f84194",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Build 1",
				Description: "note info here: fwefweewfewfwefwe",
				Uri:         "https://github4.horvatic.com/build1",
			},
			{
				Id:          "217c4d79-e5a3-4bf2-8beb-4fbd8f556c7e",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Build 2",
				Description: "wsefdfewi9v  i9 fwei9 0 scdi9 si9few",
				Uri:         "https://github4.horvatic.com/build2",
			},
			{
				Id:          "32627097-cc16-4d2e-b84e-3626b0e66021",
				ProjectId:   "785588cf-9ec0-4482-9a5f-df343cec6ac4",
				Name:        "Build 3",
				Description: "gerger  w34 tsd sfewsf se",
				Uri:         "https://github4.horvatic.com/build3",
			},
		}
	} else if id == "8c98d8dd-28a0-4d2e-896b-31b1a59fed93" {
		return []model.BuildMetaData{
			{
				Id:          "3649e08e-2b55-4959-a031-fb26b43f281a",
				ProjectId:   "8c98d8dd-28a0-4d2e-896b-31b1a59fed93",
				Name:        "Build 4",
				Description: "note info here: fwefweewfewfwefwe",
				Uri:         "https://github4.horvatic.com/build4",
			},
		}
	}
	return []model.BuildMetaData{}
}

func GetBuildsForProject(projectid string, buildId string) []model.Build {
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
