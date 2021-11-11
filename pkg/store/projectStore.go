package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/horvatic/zracni-udar-service/pkg/model"
)

type Store interface {
	GetAllProjects() []model.Project
	GetProjectById(id string) *model.Project
	CreateProject(project *model.Project) error
	UpdateProject(project *model.Project) error
	DeleteProject(projectId string) error
}

type mongoProjectStore struct {
	client     *mongo.Client
	context    context.Context
	database   string
	collection string
}

func BuildMongoProjectStore(connectionString string, database string, collection string) (Store, *mongo.Client, context.Context, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, nil, nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	return &mongoProjectStore{
		client:     client,
		database:   database,
		collection: collection,
	}, client, ctx, nil
}

func (m *mongoProjectStore) getProjectCollection() *mongo.Collection {
	db := m.client.Database(m.database)
	projects := db.Collection(m.collection)
	return projects
}

func (m *mongoProjectStore) GetAllProjects() []model.Project {
	projectsCollection := m.getProjectCollection()
	cursor, err := projectsCollection.Find(m.context, bson.M{})
	if err != nil {
		return nil
	}
	var projects []model.Project
	defer cursor.Close(m.context)
	for cursor.Next(m.context) {
		var project model.Project
		if err = cursor.Decode(&project); err != nil {
			return nil
		}
		projects = append(projects, project)
	}
	return projects
}

func (m *mongoProjectStore) GetProjectById(id string) *model.Project {
	var project model.Project
	projectsCollection := m.getProjectCollection()
	if err := projectsCollection.FindOne(m.context, bson.M{"projectId": id}).Decode(&project); err != nil {
		return nil
	}
	return &project
}

func (m *mongoProjectStore) CreateProject(project *model.Project) error {
	projectsCollection := m.getProjectCollection()
	_, err := projectsCollection.InsertOne(m.context, project)
	if err != nil {
		return err
	}
	return nil
}

func (m *mongoProjectStore) UpdateProject(project *model.Project) error {
	projectsCollection := m.getProjectCollection()
	_, err := projectsCollection.ReplaceOne(m.context, bson.M{"projectId": project.ProjectId}, project)
	if err != nil {
		return err
	}
	return nil
}

func (m *mongoProjectStore) DeleteProject(projectId string) error {
	projectsCollection := m.getProjectCollection()
	_, err := projectsCollection.DeleteOne(m.context, bson.M{"projectId": projectId})
	if err != nil {
		return err
	}
	return nil
}
