package service

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/horvatic/zracni-udar-service/pkg/model"
	"github.com/horvatic/zracni-udar-service/pkg/store"
)

type VideoService interface {
	GetVideosByProjectId(id string) ([]model.Video, ErrorType, error)
	CreateVideo(projectId string, body *io.ReadCloser) (ErrorType, error)
	UpdateVideo(videoId string, projectId string, body *io.ReadCloser) (ErrorType, error)
	DeleteVideo(videoId string, projectId string) (ErrorType, error)
	GetVideo(projectId string, videoId string) (*model.Video, ErrorType, error)
}

type videoService struct {
	store store.Store
}

func BuildVideoService(store store.Store) VideoService {
	return &videoService{
		store: store,
	}
}

func (v *videoService) GetVideosByProjectId(id string) ([]model.Video, ErrorType, error) {
	var videos []model.Video
	project := v.store.GetProjectById(id)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}
	for _, p := range project.ProjectVideos {
		videos = append(videos, model.Video{
			Id:          p.Id,
			ProjectId:   project.ProjectId,
			Name:        p.Name,
			Description: p.Description,
			Uri:         p.Uri,
		})
	}
	return videos, NoError, nil
}

func (v *videoService) CreateVideo(projectId string, body *io.ReadCloser) (ErrorType, error) {
	var video model.Video
	err := json.NewDecoder(*body).Decode(&video)
	if err != nil {
		return JsonError, err
	}
	project := v.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	project.ProjectVideos = append(project.ProjectVideos, model.ProjectVideos{
		Id:          uuid.NewString(),
		Name:        video.Name,
		Description: video.Description,
		Uri:         video.Uri,
	})
	updateErr := v.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (v *videoService) UpdateVideo(videoId string, projectId string, body *io.ReadCloser) (ErrorType, error) {
	var video model.Video
	err := json.NewDecoder(*body).Decode(&video)
	if err != nil {
		return JsonError, err
	}
	project := v.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectVideos {
		if pn.Id == videoId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find video")
	}
	project.ProjectVideos[index].Name = video.Name
	project.ProjectVideos[index].Description = video.Description
	project.ProjectVideos[index].Uri = video.Uri
	updateErr := v.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (v *videoService) DeleteVideo(videoId string, projectId string) (ErrorType, error) {
	project := v.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectVideos {
		if pn.Id == videoId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find video")
	}
	project.ProjectVideos = append(project.ProjectVideos[:index], project.ProjectVideos[index+1:]...)

	updateErr := v.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (v *videoService) GetVideo(projectId string, videoId string) (*model.Video, ErrorType, error) {
	project := v.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, p := range project.ProjectVideos {
		if p.Id == videoId {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, BadRequest, errors.New("can not find video")
	}

	return &model.Video{
		Id:          project.ProjectVideos[index].Id,
		ProjectId:   project.ProjectId,
		Name:        project.ProjectVideos[index].Name,
		Description: project.ProjectVideos[index].Description,
		Uri:         project.ProjectVideos[index].Uri,
	}, NoError, nil
}
