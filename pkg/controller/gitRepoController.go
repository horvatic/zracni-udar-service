package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horvatic/zracni-udar-service/pkg/service"
)

type GitRepoController interface {
	GetGitReposByProjectId(w http.ResponseWriter, req *http.Request)
	GetGitRepo(w http.ResponseWriter, req *http.Request)
	CreateGitRepo(w http.ResponseWriter, req *http.Request)
	UpdateGitRepo(w http.ResponseWriter, req *http.Request)
	DeleteGitRepo(w http.ResponseWriter, req *http.Request)
}

type gitRepoController struct {
	gitRepoService service.GitRepoService
}

func BuildGitRepoController(gitRepoService service.GitRepoService) GitRepoController {
	return &gitRepoController{
		gitRepoService: gitRepoService,
	}
}

func (g *gitRepoController) GetGitReposByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	blogs, errorType, err := g.gitRepoService.GetGitReposByProjectId(vars["id"])
	sendJson(w, blogs, errorType, err)
}

func (g *gitRepoController) GetGitRepo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	gitRepo, errorType, err := g.gitRepoService.GetGitRepo(vars["projectId"], vars["gitRepoId"])
	sendJson(w, gitRepo, errorType, err)
}

func (g *gitRepoController) CreateGitRepo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := g.gitRepoService.CreateGitRepo(vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (g *gitRepoController) UpdateGitRepo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := g.gitRepoService.UpdateGitRepo(vars["gitRepoId"], vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (g *gitRepoController) DeleteGitRepo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := g.gitRepoService.DeleteGitRepo(vars["gitRepoId"], vars["projectId"])
	sendCreateResult(w, errorType, err)
}
