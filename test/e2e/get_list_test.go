package e2e_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetListOK(t *testing.T) {
	c := NewClient(BaseAddr)

	name := uuid.NewString()
	listResponse, err := c.CreateList(name)
	require.NoError(t, err)

	response, err := c.GetListByID(listResponse.ID)
	require.NoError(t, err)
	assert.Equal(t, listResponse.ID, response.List.ID)
	assert.Equal(t, name, response.List.Name)
}

func TestGetListWithTasksOK(t *testing.T) {
	c := NewClient(BaseAddr)

	listName := uuid.NewString()
	listResponse, err := c.CreateList(listName)
	require.NoError(t, err)

	taskName := uuid.NewString()
	taskResponse, err := c.CreateTask(listResponse.ID, taskName)
	require.NoError(t, err)

	response, err := c.GetListByID(listResponse.ID)
	require.NoError(t, err)
	assert.Equal(t, listResponse.ID, response.List.ID)
	assert.Equal(t, listName, response.List.Name)

	assert.Equal(t, 1, len(response.List.Tasks))
	assert.Equal(t, taskResponse.ID, response.List.Tasks[0].ID)
	assert.Equal(t, taskName, response.List.Tasks[0].Name)
}

func TestGetListNotFound(t *testing.T) {
	c := NewClient(BaseAddr)

	_, err := c.GetListByID(-1)
	apiErr, ok := err.(*APIError)
	require.True(t, ok)
	assert.Equal(t, 404, apiErr.Code)
	assert.Equal(t, "list not found", apiErr.Message)
}
