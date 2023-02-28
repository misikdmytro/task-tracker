package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/misikdmytro/task-tracker/internal/database"
	"github.com/misikdmytro/task-tracker/internal/model"
	"github.com/samber/lo"
)

var (
	ErrListNotFound = fmt.Errorf("list not found")
)

type ListService interface {
	CreateList(ctx context.Context, name string) (int, error)
	GetTaskList(ctx context.Context, listID int) (model.List, error)
	AddTask(ctx context.Context, listID int, name string) (int, error)
	CloseTask(ctx context.Context, taskID int) error
}

type listService struct {
	r database.Repository
}

func NewListService(r database.Repository) ListService {
	return &listService{r: r}
}

func (l *listService) CreateList(ctx context.Context, name string) (int, error) {
	return l.r.CreateList(ctx, name)
}

func (l *listService) GetTaskList(ctx context.Context, listID int) (model.List, error) {
	result, err := l.r.GetTaskList(ctx, listID)
	if err != nil {
		return model.List{}, err
	}

	if len(result) == 0 {
		return model.List{}, ErrListNotFound
	}

	list := model.List{
		ID:        result[0].ListID,
		Name:      result[0].ListName,
		CreatedAt: result[0].ListCreatedAt,
		Tasks: lo.Map(
			lo.Filter(result, func(item model.TaskListDto, index int) bool {
				return item.TaskID != nil
			}),
			func(item model.TaskListDto, index int) model.Task {
				return model.Task{
					ID:        *item.TaskID,
					Name:      *item.TaskName,
					CreatedAt: *item.TaskCreatedAt,
				}
			}),
	}

	return list, nil
}

func (l *listService) AddTask(ctx context.Context, listID int, name string) (int, error) {
	taskID, err := l.r.CreateTask(ctx, listID, name)
	if err != nil {
		if errors.Is(err, database.ErrListForeignKeyViolation) {
			return 0, ErrListNotFound
		}

		return 0, err
	}

	return taskID, nil
}

func (l *listService) CloseTask(ctx context.Context, taskID int) error {
	return l.r.DeleteTask(ctx, taskID)
}
