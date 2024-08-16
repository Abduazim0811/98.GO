package service

import (
	"homework/internal/models"
	"homework/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService struct {
	Repo repository.TaskRepository
}

func NewTaskRepository(repo repository.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (u *TaskService) Createtask(task *models.CreateTask) error{
	return u.Repo.AddTaskMongodb(task)
}
func (u *TaskService) Gettask()([]*models.Task, error){
	return u.Repo.TaskGetMongodb()
}
func (u *TaskService) Updatetask(userID primitive.ObjectID, task models.CreateTask) error{
	return u.Repo.UpdateTaskMongodb(userID, task)
}
func (u *TaskService) Deletetask(userId primitive.ObjectID) error{
	return u.Repo.DeleteTaskMongodb(userId)
}