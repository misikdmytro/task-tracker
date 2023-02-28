package server

import (
	"net/http"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/misikdmytro/task-tracker/docs"
	"github.com/misikdmytro/task-tracker/internal/handler"
)

func NewServer(l handler.ListHandler, h handler.HealthHandler) *http.Server {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return &http.Server{
		Addr:    ":4000",
		Handler: r,
	}
}
