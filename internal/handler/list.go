package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/misikdmytro/task-tracker/internal/model"
	"github.com/misikdmytro/task-tracker/internal/service"
)

type ListHandler interface {
	CreateList(ctx *gin.Context)
	GetListByID(ctx *gin.Context)
}

type listHandler struct {
	s service.ListService
}

func NewListHandler(s service.ListService) ListHandler {
	return &listHandler{s: s}
}

func (h *listHandler) CreateList(ctx *gin.Context) {
	var request model.CreateListRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "invalid request body",
		})
		return
	}

	id, err := h.s.CreateList(ctx, request.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.CreateListResponse{
		ID: id,
	})
}

func (h *listHandler) GetListByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "id is required",
		})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "invalid id",
		})
		return
	}

	result, err := h.s.GetTaskList(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrListNotFound) {
			ctx.JSON(http.StatusNotFound, model.ErrorResponse{
				Error: "list not found",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Error: "internal server error",
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, model.GetListByIDResponse{
		List: result,
	})
}
