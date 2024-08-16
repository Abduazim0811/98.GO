package api

import (
	"homework/internal/storage"

	"github.com/gin-gonic/gin"
)

func Router() {
	handler := storage.Handler()

	r := gin.Default()

	r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks", handler.GetTask)
	r.PUT("/tasks/:id", handler.UpdateTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)
	r.Run(":7777")

}
