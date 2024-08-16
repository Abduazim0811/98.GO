package handler_test

import (
	"bytes"
	"encoding/json"
	"homework/internal/models"
	"homework/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateTask(t *testing.T) {
	r := gin.Default()
	handler := storage.Handler()
	r.POST("/tasks", handler.CreateTask)

	task := models.CreateTask{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	taskJSON, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Task created successfully", response["message"])
}

func TestGetTasks(t *testing.T) {
	r := gin.Default()
	handler := storage.Handler()
	r.GET("/tasks", handler.GetTask)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string][]models.Task
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response["Tasks"])
}

// func TestUpdateTask(t *testing.T) {
// 	r := gin.Default()
// 	handler := storage.Handler()
// 	r.PUT("/tasks/:id", handler.UpdateTask)

// 	taskID := primitive.NewObjectID().Hex()

// 	updatedTask := models.CreateTask{
// 		Title:       "Updated Task",
// 		Description: "This task has been updated",
// 	}
// 	taskJSON, _ := json.Marshal(updatedTask)

// 	req, _ := http.NewRequest("PUT", "/tasks/"+taskID, bytes.NewBuffer(taskJSON))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	var response map[string]string
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Task updated successfully", response["message"])
// }

func TestDeleteTask(t *testing.T) {
	r := gin.Default()
	handler := storage.Handler()
	r.DELETE("/tasks/:id", handler.DeleteTask)

	taskID := primitive.NewObjectID().Hex()

	req, _ := http.NewRequest("DELETE", "/tasks/"+taskID, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Task deleted successfully", response["message"])
}
