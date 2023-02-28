package e2e_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/misikdmytro/task-tracker/pkg/model"
)

type Client interface {
	CreateList(string) (model.CreateListResponse, error)
	GetListByID(int) (model.GetListByIDResponse, error)
	CreateTask(int, string) (model.AddTaskResponse, error)
	CloseTask(int) error
}

type client struct {
	c           http.Client
	baseAddress string
}

type APIError struct {
	Message string
	Code    int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("unexpected status code: %d. details: %s", e.Code, e.Message)
}

func NewClient(baseAddress string) Client {
	return &client{baseAddress: baseAddress}
}

func (c *client) CloseTask(taskID int) error {
	request, err := http.NewRequest(http.MethodDelete, c.baseAddress+"/tasks/"+fmt.Sprint(taskID), nil)
	if err != nil {
		return err
	}

	response, err := c.c.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		var result model.ErrorResponse
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			return err
		}

		return &APIError{
			Message: result.Error,
			Code:    response.StatusCode,
		}
	}

	return nil
}

func (c *client) CreateList(name string) (model.CreateListResponse, error) {
	m := model.CreateListRequest{
		Name: name,
	}

	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return model.CreateListResponse{}, err
	}

	request, err := http.NewRequest(http.MethodPut, c.baseAddress+"/lists/", bytes.NewReader(jsonBytes))
	if err != nil {
		return model.CreateListResponse{}, err
	}

	response, err := c.c.Do(request)
	if err != nil {
		return model.CreateListResponse{}, err
	}

	if response.StatusCode != http.StatusCreated {
		var result model.ErrorResponse
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			return model.CreateListResponse{}, err
		}

		return model.CreateListResponse{}, &APIError{
			Message: result.Error,
			Code:    response.StatusCode,
		}
	}

	var result model.CreateListResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return model.CreateListResponse{}, err
	}

	return result, nil
}

func (c *client) CreateTask(listID int, name string) (model.AddTaskResponse, error) {
	m := model.AddTaskRequest{
		Name: name,
	}

	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return model.AddTaskResponse{}, err
	}

	request, err := http.NewRequest(http.MethodPut, c.baseAddress+"/lists/"+fmt.Sprint(listID)+"/tasks", bytes.NewReader(jsonBytes))
	if err != nil {
		return model.AddTaskResponse{}, err
	}

	response, err := c.c.Do(request)
	if err != nil {
		return model.AddTaskResponse{}, err
	}

	if response.StatusCode != http.StatusCreated {
		var result model.ErrorResponse
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			return model.AddTaskResponse{}, err
		}

		return model.AddTaskResponse{}, &APIError{
			Message: result.Error,
			Code:    response.StatusCode,
		}
	}

	var result model.AddTaskResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return model.AddTaskResponse{}, err
	}

	return result, nil
}

func (c *client) GetListByID(taskID int) (model.GetListByIDResponse, error) {
	request, err := http.NewRequest(http.MethodGet, c.baseAddress+"/lists/"+fmt.Sprint(taskID), nil)
	if err != nil {
		return model.GetListByIDResponse{}, err
	}

	response, err := c.c.Do(request)
	if err != nil {
		return model.GetListByIDResponse{}, err
	}

	if response.StatusCode != http.StatusOK {
		var result model.ErrorResponse
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			return model.GetListByIDResponse{}, err
		}

		return model.GetListByIDResponse{}, &APIError{
			Message: result.Error,
			Code:    response.StatusCode,
		}
	}

	var result model.GetListByIDResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return model.GetListByIDResponse{}, err
	}

	return result, nil
}
