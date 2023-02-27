package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/misikdmytro/task-tracker/internal/database"
	"github.com/misikdmytro/task-tracker/internal/model"
	"github.com/misikdmytro/task-tracker/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
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

func TestGetTaskList(t *testing.T) {
	input := []struct {
		name          string
		dbResult      []model.TaskListDto
		dbError       error
		expected      model.List
		expectedError error
	}{
		{
			"success",
			[]model.TaskListDto{
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
			},
			nil,
			model.List{
				ID:        1,
				Name:      "list1",
				CreatedAt: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC),
				Tasks: []model.Task{
					{
						ID:        1,
						Name:      "task1",
						CreatedAt: time.Date(2021, 2, 3, 4, 5, 6, 7, time.UTC),
					},
					{
						ID:        2,
						Name:      "task2",
						CreatedAt: time.Date(2021, 3, 4, 5, 6, 7, 8, time.UTC),
					},
				},
			},
			nil,
		},
		{
			"success without tasks",
			[]model.TaskListDto{
				{
					ListID:        1,
					ListName:      "list1",
					ListCreatedAt: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC),
				},
			},
			nil,
			model.List{
				ID:        1,
				Name:      "list1",
				CreatedAt: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC),
				Tasks:     []model.Task{},
			},
			nil,
		},
		{
			"list not found error #1",
			[]model.TaskListDto{},
			nil,
			model.List{},
			service.ErrListNotFound,
		},
		{
			"list not found error #2",
			nil,
			nil,
			model.List{},
			service.ErrListNotFound,
		},
	}

	for _, tc := range input {
		t.Run(tc.name, func(t *testing.T) {
			r := &repositoryMock{}
			s := service.NewListService(r)

			r.On("GetTaskList", mock.Anything, 1).Return(tc.dbResult, tc.dbError)

			res, err := s.GetTaskList(context.Background(), 1)

			assert.Equal(t, tc.expected, res)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestCreateList(t *testing.T) {
	r := &repositoryMock{}
	s := service.NewListService(r)

	r.On("CreateList", mock.Anything, "list1").Return(1, nil)

	s.CreateList(context.Background(), "list1")

	r.AssertCalled(t, "CreateList", mock.Anything, "list1")
	r.AssertNumberOfCalls(t, "CreateList", 1)
}
