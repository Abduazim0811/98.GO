package storage

import (
	"context"
	"homework/internal/http/api/handler"
	"homework/internal/repository/mongodb"
	"homework/internal/service"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewTask() (*mongo.Client, *mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("Tasks").Collection("task")
	return client, collection, nil
}

func Handler() *handler.TaskHandler {
	client, collection, err := NewTask()
	if err != nil {
		log.Println("connection mongodb error")
	}

	repo := mongodb.NewTaskMongoDb(client, collection)

	service := service.NewTaskRepository(repo)

	handler := handler.NewTaskHandler(service)
	return handler
}
