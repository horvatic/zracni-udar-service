package service

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/horvatic/zracni-udar-service/pkg/model"
	"github.com/horvatic/zracni-udar-service/pkg/store"
)

type BlogService interface {
	GetBlogsByProjectId(id string) ([]model.Blog, ErrorType, error)
	CreateBlog(projectId string, body *io.ReadCloser) (ErrorType, error)
	UpdateBlog(blogId string, projectId string, body *io.ReadCloser) (ErrorType, error)
	DeleteBlog(blogId string, projectId string) (ErrorType, error)
	GetBlog(projectId string, blogId string) (*model.Blog, ErrorType, error)
}

type blogService struct {
	store store.Store
}

func BuildBlogService(store store.Store) BlogService {
	return &blogService{
		store: store,
	}
}

func (b *blogService) GetBlogsByProjectId(id string) ([]model.Blog, ErrorType, error) {
	var blogs []model.Blog
	project := b.store.GetProjectById(id)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}

	for _, p := range project.ProjectBlogs {
		blogs = append(blogs, model.Blog{
			Id:          p.Id,
			ProjectId:   project.ProjectId,
			Name:        p.Name,
			Description: p.Description,
			Uri:         p.Uri,
		})
	}
	return blogs, NoError, nil
}

func (b *blogService) CreateBlog(projectId string, body *io.ReadCloser) (ErrorType, error) {
	var blog model.Blog
	err := json.NewDecoder(*body).Decode(&blog)
	if err != nil {
		return JsonError, err
	}
	project := b.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	project.ProjectBlogs = append(project.ProjectBlogs, model.ProjectBlogs{
		Id:          uuid.NewString(),
		Name:        blog.Name,
		Description: blog.Description,
		Uri:         blog.Uri,
	})
	updateErr := b.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (b *blogService) UpdateBlog(blogId string, projectId string, body *io.ReadCloser) (ErrorType, error) {
	var blog model.Blog
	err := json.NewDecoder(*body).Decode(&blog)
	if err != nil {
		return JsonError, err
	}
	project := b.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectBlogs {
		if pn.Id == blogId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find blog")
	}
	project.ProjectBlogs[index].Name = blog.Name
	project.ProjectBlogs[index].Description = blog.Description
	project.ProjectBlogs[index].Uri = blog.Uri
	updateErr := b.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (b *blogService) DeleteBlog(blogId string, projectId string) (ErrorType, error) {
	project := b.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectBlogs {
		if pn.Id == blogId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find blog")
	}
	project.ProjectBlogs = append(project.ProjectBlogs[:index], project.ProjectBlogs[index+1:]...)

	updateErr := b.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (b *blogService) GetBlog(projectId string, blogId string) (*model.Blog, ErrorType, error) {
	project := b.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, p := range project.ProjectBlogs {
		if p.Id == blogId {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, BadRequest, errors.New("can not find blog")
	}

	return &model.Blog{
		Id:          project.ProjectBlogs[index].Id,
		ProjectId:   project.ProjectId,
		Name:        project.ProjectBlogs[index].Name,
		Description: project.ProjectBlogs[index].Description,
		Uri:         project.ProjectBlogs[index].Uri,
	}, NoError, nil
}
