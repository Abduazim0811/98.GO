package mongodb

import (
	"context"
	"log"
	"time"

	"homework/internal/models"
	"homework/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskMongodb struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewTaskMongoDb(client *mongo.Client, collection *mongo.Collection) repository.TaskRepository {
	return &TaskMongodb{client: client, collection: collection}
}

func (u *TaskMongodb) AddTaskMongodb(task *models.CreateTask) error {
	var Task models.Task

	Task.Title = task.Title
	Task.Description = task.Description
	Task.Status = "created"
	Task.CreatedAt = time.Now()
	Task.UpdatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := u.collection.InsertOne(ctx, Task)
	if err != nil {
		log.Println("Error adding task:", err)
		return err
	}
	return nil
}

func (u *TaskMongodb) TaskGetMongodb() ([]*models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var Tasks []*models.Task
	cursor, err := u.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error fetching Tasks:", err)
		return nil, err
	}

	if err = cursor.All(ctx, &Tasks); err != nil {
		log.Println("Error decoding Tasks:", err)
		return nil, err
	}
	return Tasks, nil
}

func (u *TaskMongodb) UpdateTaskMongodb(TaskID primitive.ObjectID, task models.CreateTask) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": TaskID}
	update := bson.M{"$set": bson.M{"title": task.Title, "description": task.Description, "status": "updated", "updated_at": time.Now()}}
	var UpdatedTask models.Task
	err := u.collection.FindOneAndUpdate(ctx, filter, update).Decode(&UpdatedTask)
	if err != nil {
		log.Println("Error updating Task:", err)
		return err
	}
	return nil
}

func (u *TaskMongodb) DeleteTaskMongodb(TaskID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// res := u.collection.FindOne(ctx, bson.M{"_id": TaskID})
	// if res.Err() != nil {
	// 	if res.Err() == mongo.ErrNoDocuments {
	// 		return fmt.Errorf("task not found")
	// 	}
	// 	return res.Err()
	// }
	_, err := u.collection.DeleteOne(ctx, bson.M{"_id": TaskID})
	if err != nil {
		log.Println("Error deleting Task:", err)
		return err
	}
	return nil
}
