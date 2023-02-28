package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/misikdmytro/task-tracker/internal/handler"
)

func NewServer(h handler.ListHandler) *http.Server {
	r := gin.Default()

	lists := r.Group("/lists")
	{
		lists.GET("/:id", h.GetListByID)
		lists.PUT("", h.CreateList)
		lists.PUT("/:id/tasks", h.AddTask)
	}

	tasks := r.Group("/tasks")
	{
		tasks.DELETE("/:id", h.CloseTask)
	}

	return &http.Server{
		Addr:    ":4000",
		Handler: r,
	}
}
