package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horvatic/zracni-udar-service/pkg/service"
)

type VideoController interface {
	GetVideosByProjectId(w http.ResponseWriter, req *http.Request)
	GetVideo(w http.ResponseWriter, req *http.Request)
	CreateVideo(w http.ResponseWriter, req *http.Request)
	UpdateVideo(w http.ResponseWriter, req *http.Request)
	DeleteVideo(w http.ResponseWriter, req *http.Request)
}

type videoController struct {
	videoService service.VideoService
}

func BuildVideoController(videoService service.VideoService) VideoController {
	return &videoController{
		videoService: videoService,
	}
}

func (v *videoController) GetVideosByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	videos, errorType, err := v.videoService.GetVideosByProjectId(vars["id"])
	sendJson(w, videos, errorType, err)
}

func (v *videoController) GetVideo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	video, errorType, err := v.videoService.GetVideo(vars["projectId"], vars["videoId"])
	sendJson(w, video, errorType, err)
}

func (v *videoController) CreateVideo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := v.videoService.CreateVideo(vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (v *videoController) UpdateVideo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := v.videoService.UpdateVideo(vars["videoId"], vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (v *videoController) DeleteVideo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := v.videoService.DeleteVideo(vars["videoId"], vars["projectId"])
	sendCreateResult(w, errorType, err)
}
