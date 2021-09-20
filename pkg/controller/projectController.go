package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horvatic/zracni-udar-service/pkg/service"
)

func GetAllProjectsMetaData(w http.ResponseWriter, req *http.Request) {
	projectMetaData := service.GetProjectsMetaData()
	sendJson(w, projectMetaData)
}

func GetProjectById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	projects := service.GetProjectById(vars["id"])
	sendJson(w, projects)
}

func GetNotesByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	notes := service.GetNotesByProjectId(vars["id"])
	sendJson(w, notes)
}

func GetBlogsByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	blogs := service.GetBlogsByProjectId(vars["id"])
	sendJson(w, blogs)
}

func GetVideosByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	videos := service.GetVideosByProjectId(vars["id"])
	sendJson(w, videos)
}

func GetDiagramsByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	videos := service.GetDiagramsByProjectId(vars["id"])
	sendJson(w, videos)
}

func GetGitReposByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	gitRepos := service.GetGitReposByProjectId(vars["id"])
	sendJson(w, gitRepos)
}

func GetBuildMetaDatasByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	builds := service.GetBuildMetaDatasByProjectId(vars["id"])
	sendJson(w, builds)
}

func GetBuildsForProject(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	deployments := service.GetBuildsForProject(vars["projectId"], vars["buildId"])
	sendJson(w, deployments)
}

func sendJson(w http.ResponseWriter, j interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(j)
}
