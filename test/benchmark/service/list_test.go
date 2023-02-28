package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/misikdmytro/task-tracker/internal/database"
	"github.com/misikdmytro/task-tracker/internal/model"
	"github.com/misikdmytro/task-tracker/internal/service"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

// Ping implements database.Repository
func (r *repositoryMock) Ping(ctx context.Context) error {
	args := r.Called(ctx)
	return args.Error(0)
}

func (r *repositoryMock) DeleteTask(ctx context.Context, taskID int) error {
	args := r.Called(ctx, taskID)
	return args.Error(0)
}

func (r *repositoryMock) CreateTask(ctx context.Context, listID int, name string) (int, error) {
	args := r.Called(ctx, listID, name)
	return args.Int(0), args.Error(1)
}

func (r *repositoryMock) CreateList(ctx context.Context, name string) (int, error) {
	args := r.Called(ctx, name)
	return args.Int(0), args.Error(1)
}

func (r *repositoryMock) GetTaskList(ctx context.Context, listID int) ([]model.TaskListDto, error) {
	args := r.Called(ctx, listID)
	return args.Get(0).([]model.TaskListDto), args.Error(1)
}

var _ database.Repository = (*repositoryMock)(nil)

func pointer[T any](v T) *T {
	return &v
}

func BenchmarkGetTaskList(b *testing.B) {
	r := repositoryMock{}
	s := service.NewListService(&r)

	r.On("GetTaskList", mock.Anything, mock.Anything).Return([]model.TaskListDto{
		{
			ListID:        1,
			ListName:      "list1",
			ListCreatedAt: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC),
			TaskID:        pointer(1),
			TaskName:      pointer("task1"),
			TaskCreatedAt: pointer(time.Date(2021, 2, 3, 4, 5, 6, 7, time.UTC)),
		},
		{
			ListID:        1,
			ListName:      "list1",
			ListCreatedAt: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC),
			TaskID:        pointer(2),
			TaskName:      pointer("task2"),
			TaskCreatedAt: pointer(time.Date(2021, 3, 4, 5, 6, 7, 8, time.UTC)),
		},
		{
			ListID:        1,
			ListName:      "list1",
			ListCreatedAt: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC),
			TaskID:        pointer(3),
			TaskName:      pointer("task3"),
			TaskCreatedAt: pointer(time.Date(2021, 4, 5, 6, 7, 8, 9, time.UTC)),
		},
		{
			ListID:        1,
			ListName:      "list1",
			ListCreatedAt: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC),
			TaskID:        pointer(4),
			TaskName:      pointer("task4"),
			TaskCreatedAt: pointer(time.Date(2021, 5, 6, 7, 8, 9, 10, time.UTC)),
		},
	}, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.GetTaskList(context.Background(), 1)
	}
}
