package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/misikdmytro/task-tracker/internal/handler"
)

func NewServer(l handler.ListHandler, h handler.HealthHandler) *http.Server {
	r := gin.Default()

	lists := r.Group("/lists")
	{
		lists.GET("/:id", l.GetListByID)
		lists.PUT("", l.CreateList)
		lists.PUT("/:id/tasks", l.AddTask)
	}

	tasks := r.Group("/tasks")
	{
		tasks.DELETE("/:id", l.CloseTask)
	}

	health := r.Group("/health")
	{
		health.GET("", h.HealthCheck)
	}

	return &http.Server{
		Addr:    ":4000",
		Handler: r,
	}
}
