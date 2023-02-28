package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	inmodel "github.com/misikdmytro/task-tracker/internal/model"
	"github.com/misikdmytro/task-tracker/internal/service"
	"github.com/misikdmytro/task-tracker/pkg/model"
	"github.com/samber/lo"
)

type ListHandler interface {
	CreateList(ctx *gin.Context)
	GetListByID(ctx *gin.Context)
	AddTask(ctx *gin.Context)
	CloseTask(ctx *gin.Context)
}

type listHandler struct {
	s service.ListService
}

func NewListHandler(s service.ListService) ListHandler {
	return &listHandler{s: s}
}

// CreateList godoc
//
//	@Summary		create list
//	@Description	create list
//	@Tags			lists
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.CreateListRequest	true	"request body"
//	@Success		201		{object}	model.CreateListResponse
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		500		{object}	model.ErrorResponse
//	@Router			/lists [put]
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

// GetListByID godoc
//
//	@Summary		get list by id
//	@Description	get list by id
//	@Tags			lists
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"list id"
//	@Success		200	{object}	model.GetListByIDResponse
//	@Failure		400	{object}	model.ErrorResponse
//	@Failure		404	{object}	model.ErrorResponse
//	@Failure		500	{object}	model.ErrorResponse
//	@Router			/lists/{id} [get]
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
		List: model.List{
			ID:        result.ID,
			Name:      result.Name,
			CreatedAt: result.CreatedAt,
			Tasks: lo.Map(result.Tasks, func(task inmodel.Task, index int) model.Task {
				return model.Task{
					ID:        task.ID,
					Name:      task.Name,
					CreatedAt: task.CreatedAt,
				}
			}),
		},
	})
}

// AddTask godoc
//
//	@Summary		add task to list
//	@Description	add task to list
//	@Tags			lists
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"list id"
//	@Param			request	body		model.AddTaskRequest	true	"request body"
//	@Success		201		{object}	model.AddTaskResponse
//	@Failure		400		{object}	model.ErrorResponse
//	@Failure		404		{object}	model.ErrorResponse
//	@Failure		500		{object}	model.ErrorResponse
//	@Router			/lists/{id}/tasks [put]
func (h *listHandler) AddTask(ctx *gin.Context) {
	var request model.AddTaskRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "invalid request body",
		})
		return
	}

	listIDParam := ctx.Param("id")
	if listIDParam == "" {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "id is required",
		})
		return
	}

	listID, err := strconv.Atoi(listIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "invalid id",
		})
		return
	}

	id, err := h.s.AddTask(ctx, listID, request.Name)
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

	ctx.JSON(http.StatusCreated, model.AddTaskResponse{
		ID: id,
	})
}

// CloseTask godoc
//
//	@Summary		close task
//	@Description	close task
//	@Tags			lists
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"task id"
//	@Success		204
//	@Failure		400	{object}	model.ErrorResponse
//	@Failure		500	{object}	model.ErrorResponse
//	@Router			/tasks/{id} [delete]
func (h *listHandler) CloseTask(ctx *gin.Context) {
	taskIDParam := ctx.Param("id")
	if taskIDParam == "" {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "id is required",
		})
		return
	}

	taskID, err := strconv.Atoi(taskIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "invalid id",
		})
		return
	}

	err = h.s.CloseTask(ctx, taskID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "internal server error",
		})

		return
	}

	ctx.Status(http.StatusNoContent)
}
