package repository

import (
	"homework/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepository interface {
	AddTaskMongodb(task *models.CreateTask) error
	TaskGetMongodb() ([]*models.Task, error)
	UpdateTaskMongodb(taskID primitive.ObjectID, task models.CreateTask) error
	DeleteTaskMongodb(taskID primitive.ObjectID) error
}
