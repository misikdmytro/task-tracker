package e2e_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCloseTaskNoContent(t *testing.T) {
	c := NewClient(BaseAddr)

	listResponse, err := c.CreateList(uuid.NewString())
	require.NoError(t, err)

	taskResponse, err := c.CreateTask(listResponse.ID, uuid.NewString())
	require.NoError(t, err)

	err = c.CloseTask(taskResponse.ID)
	require.NoError(t, err)
}

func TestCloseTaskNoTaskNoContent(t *testing.T) {
	c := NewClient(BaseAddr)

	err := c.CloseTask(-1)
	require.NoError(t, err)
}
