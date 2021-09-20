package store

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/horvatic/zracni-udar-service/pkg/model"
)

type Store interface {
	GetProjectsMetaData() []model.ProjectMetaData
	GetProjectById(id string) model.ProjectMetaData
	GetNotesByProjectId(id string) []model.Note
	GetBlogsByProjectId(id string) []model.Blog
	GetVideosByProjectId(id string) []model.Video
	GetDiagramsByProjectId(id string) []model.Diagram
	GetGitReposByProjectId(id string) []model.GitRepo
	GetBuildMetaDatasByProjectId(id string) []model.BuildMetaData
}

type mongoProjectStore struct {
	connectionString string
	database         string
	collection       string
}

type connection struct {
	client  *mongo.Client
	context context.Context
}

func BuildMongoProjectStore(connectionString string, database string, collection string) Store {
	return &mongoProjectStore{
		connectionString: connectionString,
		database:         database,
		collection:       collection,
	}
}

func (m *mongoProjectStore) connect() (*connection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.connectionString))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &connection{
		client:  client,
		context: ctx,
	}, nil
}

func (m *mongoProjectStore) getProjectCollection(connection *connection) *mongo.Collection {
	db := connection.client.Database(m.database)
	projects := db.Collection(m.collection)
	return projects
}

func (m *mongoProjectStore) disconnect(connection *connection) {
	connection.client.Disconnect(connection.context)
}

func (m *mongoProjectStore) getProject(id string) *model.Project {
	conn, err := m.connect()
	if err != nil {
		return &model.Project{}
	}
	defer m.disconnect(conn)

	var project model.Project
	projectsCollection := m.getProjectCollection(conn)
	if err = projectsCollection.FindOne(conn.context, bson.M{"projectId": id}).Decode(&project); err != nil {
		log.Fatal(err)
	}
	return &project
}

func (m *mongoProjectStore) GetProjectsMetaData() []model.ProjectMetaData {
	conn, err := m.connect()
	if err != nil {
		return []model.ProjectMetaData{}
	}
	defer m.disconnect(conn)

	projectsCollection := m.getProjectCollection(conn)
	cursor, err := projectsCollection.Find(conn.context, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var projects []model.ProjectMetaData
	defer cursor.Close(conn.context)
	for cursor.Next(conn.context) {
		var project model.Project
		if err = cursor.Decode(&project); err != nil {
			log.Fatal(err)
		}
		projects = append(projects, model.ProjectMetaData{
			Id:          project.ProjectId,
			Name:        project.Name,
			Description: project.Description,
		})
	}

	return projects
}

func (m *mongoProjectStore) GetProjectById(id string) model.ProjectMetaData {
	project := m.getProject(id)

	return model.ProjectMetaData{
		Id:          project.ProjectId,
		Name:        project.Name,
		Description: project.Description,
	}
}

func (m *mongoProjectStore) GetNotesByProjectId(id string) []model.Note {
	var notes []model.Note
	project := m.getProject(id)

	for _, p := range project.ProjectNotes {
		notes = append(notes, model.Note{
			Id:        p.Id,
			ProjectId: project.ProjectId,
			Name:      p.Name,
			Note:      p.Note,
		})
	}

	return notes
}

func (m *mongoProjectStore) GetBlogsByProjectId(id string) []model.Blog {
	var blogs []model.Blog
	project := m.getProject(id)
	for _, p := range project.ProjectBlogs {
		blogs = append(blogs, model.Blog{
			Id:          p.Id,
			ProjectId:   project.ProjectId,
			Name:        p.Name,
			Description: p.Description,
			Uri:         p.Uri,
		})
	}
	return blogs
}

func (m *mongoProjectStore) GetVideosByProjectId(id string) []model.Video {
	var videos []model.Video
	project := m.getProject(id)
	for _, p := range project.ProjectVideos {
		videos = append(videos, model.Video{
			Id:          p.Id,
			ProjectId:   project.ProjectId,
			Name:        p.Name,
			Description: p.Description,
			Uri:         p.Uri,
		})
	}
	return videos
}

func (m *mongoProjectStore) GetDiagramsByProjectId(id string) []model.Diagram {
	var diagrams []model.Diagram
	project := m.getProject(id)
	for _, p := range project.ProjectDiagrams {
		diagrams = append(diagrams, model.Diagram{
			Id:          p.Id,
			ProjectId:   project.ProjectId,
			Name:        p.Name,
			Description: p.Description,
			Uri:         p.Uri,
		})
	}
	return diagrams
}

func (m *mongoProjectStore) GetGitReposByProjectId(id string) []model.GitRepo {
	var gitRepo []model.GitRepo
	project := m.getProject(id)
	for _, p := range project.ProjectGitRepos {
		gitRepo = append(gitRepo, model.GitRepo{
			Id:          p.Id,
			ProjectId:   project.ProjectId,
			Name:        p.Name,
			Description: p.Description,
			Uri:         p.Uri,
		})
	}
	return gitRepo
}

func (m *mongoProjectStore) GetBuildMetaDatasByProjectId(id string) []model.BuildMetaData {
	var buildMetaData []model.BuildMetaData
	project := m.getProject(id)
	for _, p := range project.ProjectBuildsMetaData {
		buildMetaData = append(buildMetaData, model.BuildMetaData{
			Id:          p.Id,
			ProjectId:   project.ProjectId,
			Name:        p.Name,
			Description: p.Description,
			Uri:         p.Uri,
		})
	}
	return buildMetaData
}
