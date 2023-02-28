package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/misikdmytro/task-tracker/internal/service"
	"github.com/misikdmytro/task-tracker/pkg/model"
)

type HealthHandler interface {
	HealthCheck(ctx *gin.Context)
}

type healthHandler struct {
	s service.HealthService
}

func NewHealthHandler(s service.HealthService) HealthHandler {
	return &healthHandler{s: s}
}

// HealthCheck godoc
//
//	@Summary		health check
//	@Description	health check
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.HealthResponse
//	@Failure		500	{object}	model.ErrorResponse
//	@Router			/health [get]
func (h *healthHandler) HealthCheck(ctx *gin.Context) {
	if err := h.s.HealthCheck(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.HealthResponse{
		Status: "OK",
	})
}
