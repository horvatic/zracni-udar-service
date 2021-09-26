package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horvatic/zracni-udar-service/pkg/service"
)

type BlogController interface {
	GetBlogsByProjectId(w http.ResponseWriter, req *http.Request)
	GetBlog(w http.ResponseWriter, req *http.Request)
	CreateBlog(w http.ResponseWriter, req *http.Request)
	UpdateBlog(w http.ResponseWriter, req *http.Request)
	DeleteBlog(w http.ResponseWriter, req *http.Request)
}

type blogController struct {
	blogService service.BlogService
}

func BuildBlogController(blogService service.BlogService) BlogController {
	return &blogController{
		blogService: blogService,
	}
}

func (b *blogController) GetBlogsByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	blogs, errorType, err := b.blogService.GetBlogsByProjectId(vars["id"])
	sendJson(w, blogs, errorType, err)
}

func (b *blogController) GetBlog(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	blog, errorType, err := b.blogService.GetBlog(vars["projectId"], vars["blogId"])
	sendJson(w, blog, errorType, err)
}

func (b *blogController) CreateBlog(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := b.blogService.CreateBlog(vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (b *blogController) UpdateBlog(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := b.blogService.UpdateBlog(vars["blogId"], vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (b *blogController) DeleteBlog(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := b.blogService.DeleteBlog(vars["blogId"], vars["projectId"])
	sendCreateResult(w, errorType, err)
}
