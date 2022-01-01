package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
        "go.mongodb.org/mongo-driver/mongo"
        "go.mongodb.org/mongo-driver/mongo/options"

	"github.com/google/uuid"
)

type Deployment struct {
        ID             primitive.ObjectID `bson:"_id,omitempty"`
        DeploymentId   string             `bson:"deploymentId"`
        RepoUri        string             `bson:"repoUri"`
        DeploymentName string             `bson:"deploymentName"`
        Env            string             `bson:"env"`
        TargetScript   string             `bson:"targetScript"`
        GitSha         string             `bson:"gitSha"`
 	BranchName     string             `bson:"branchName"`
	GitShortSha    string		  `bson:"gitShortSha"`
}

func main() {
	fmt.Println("Starting Deployment")
	connectionString  := os.Getenv("DEPLOYMENT_MONGO_CONNECTION_STRING")
        database := os.Getenv("DEPLOYMENT_MONGO_DATABASE")
	collection := os.Getenv("DEPLOYMENT_MONGO_COLLECTION")
	deploymentEnv := os.Getenv("DEPLOYMENT_ENV")
	gitSha := os.Getenv("DEPLOYMENT_GITSHA")
	branchName := os.Getenv("DEPLOYMENT_BRANCH_NAME")
	gitShortSha := os.Getenv("DEPLOYMENT_GIT_SHORT_SHA")


	fmt.Println("Setting Up Connection To DB")
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	db := client.Database(database)
	deploymentCollection := db.Collection(collection)

	fmt.Println("Connection Setup")
	deploymentRequest := &Deployment {
		DeploymentId: uuid.NewString(),
		RepoUri: "https://github.com/horvatic/zracni-udar-service.git",
		DeploymentName: "zracni-udar-service",
		Env: deploymentEnv,
		TargetScript: "k8deploy.sh",
		GitSha: gitSha,
		BranchName: branchName,
		GitShortSha: gitShortSha,
	}

	_, err = deploymentCollection.InsertOne(ctx, deploymentRequest)
	if err != nil {
        	fmt.Println(err)
                return
        }

        client.Disconnect(ctx)

	return
}
