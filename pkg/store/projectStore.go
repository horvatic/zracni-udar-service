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
	GetAllProjects() []model.Project
	GetProjectById(id string) *model.Project
	CreateProject(project *model.Project) error
	UpdateProject(project *model.Project) error
	DeleteProject(projectId string) error
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

func (m *mongoProjectStore) GetAllProjects() []model.Project {
	conn, err := m.connect()
	if err != nil {
		return []model.Project{}
	}
	defer m.disconnect(conn)

	projectsCollection := m.getProjectCollection(conn)
	cursor, err := projectsCollection.Find(conn.context, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var projects []model.Project
	defer cursor.Close(conn.context)
	for cursor.Next(conn.context) {
		var project model.Project
		if err = cursor.Decode(&project); err != nil {
			log.Fatal(err)
		}
		projects = append(projects, project)
	}

	return projects
}

func (m *mongoProjectStore) GetProjectById(id string) *model.Project {
	conn, err := m.connect()
	if err != nil {
		return &model.Project{}
	}
	defer m.disconnect(conn)

	var project model.Project
	projectsCollection := m.getProjectCollection(conn)
	if err = projectsCollection.FindOne(conn.context, bson.M{"projectId": id}).Decode(&project); err != nil {
		log.Fatal(err)
		return nil
	}
	return &project
}

func (m *mongoProjectStore) CreateProject(project *model.Project) error {
	conn, conErr := m.connect()
	if conErr != nil {
		return conErr
	}
	defer m.disconnect(conn)
	projectsCollection := m.getProjectCollection(conn)
	_, err := projectsCollection.InsertOne(conn.context, project)
	if err != nil {
		return err
	}
	return nil
}

func (m *mongoProjectStore) UpdateProject(project *model.Project) error {
	conn, conErr := m.connect()
	if conErr != nil {
		return conErr
	}
	defer m.disconnect(conn)
	projectsCollection := m.getProjectCollection(conn)
	_, err := projectsCollection.ReplaceOne(conn.context, bson.M{"projectId": project.ProjectId}, project)
	if err != nil {
		return err
	}
	return nil
}

func (m *mongoProjectStore) DeleteProject(projectId string) error {
	conn, conErr := m.connect()
	if conErr != nil {
		return conErr
	}
	defer m.disconnect(conn)
	projectsCollection := m.getProjectCollection(conn)
	_, err := projectsCollection.DeleteOne(conn.context, bson.M{"projectId": projectId})
	if err != nil {
		return err
	}
	return nil
}
