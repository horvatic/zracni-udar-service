package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID                    primitive.ObjectID      `bson:"_id,omitempty"`
	ProjectId             string                  `bson:"projectId"`
	Name                  string                  `bson:"name"`
	Description           string                  `bson:"description"`
	ProjectNotes          []ProjectNotes          `bson:"notes"`
	ProjectBlogs          []ProjectBlogs          `bson:"blogs"`
	ProjectVideos         []ProjectVideos         `bson:"videos"`
	ProjectDiagrams       []ProjectDiagrams       `bson:"diagrams"`
	ProjectGitRepos       []ProjectGitRepos       `bson:"git_repos"`
	ProjectBuildsMetaData []ProjectBuildsMetaData `bson:"builds"`
}

type ProjectNotes struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
	Note string `bson:"note"`
}

type ProjectBlogs struct {
	Id          string `bson:"id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Uri         string `bson:"uri"`
}

type ProjectVideos struct {
	Id          string `bson:"id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Uri         string `bson:"uri"`
}

type ProjectDiagrams struct {
	Id          string `bson:"id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Uri         string `bson:"uri"`
}

type ProjectGitRepos struct {
	Id          string `bson:"id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Uri         string `bson:"uri"`
}

type ProjectBuildsMetaData struct {
	Id          string `bson:"id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Uri         string `bson:"uri"`
	RepoOwner   string `bson:"repo_owner"`
	RepoName    string `bson:"repo_name"`
}
