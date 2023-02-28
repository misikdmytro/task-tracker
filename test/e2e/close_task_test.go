package e2e_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCloseTaskNoContent(t *testing.T) {
	s, start, close := Setup(t)
	defer close()
	start()

	db, err := s.F.NewDB()
	require.NoError(t, err)
	defer db.Close()

	var listID int
	require.NoError(t, db.Get(&listID, "INSERT INTO tbl_lists (name) VALUES ($1) RETURNING id", uuid.NewString()))

	var taskID int
	require.NoError(t, db.Get(&taskID, "INSERT INTO tbl_tasks (name, list_id) VALUES ($1, $2) RETURNING id", uuid.NewString(), listID))

	request, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("http://localhost:4000/tasks/%d", taskID),
		nil,
	)
	require.NoError(t, err)

	client := http.Client{}
	response, err := client.Do(request)
	require.NoError(t, err)

	require.Equal(t, http.StatusNoContent, response.StatusCode)
}

func TestCloseTaskNoTaskNoContent(t *testing.T) {
	s, start, close := Setup(t)
	defer close()
	start()

	db, err := s.F.NewDB()
	require.NoError(t, err)
	defer db.Close()

	var taskID int
	require.NoError(t, db.Get(&taskID, "SELECT MAX(id) + 1 FROM tbl_tasks"))

	request, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("http://localhost:4000/tasks/%d", taskID),
		nil,
	)
	require.NoError(t, err)

	client := http.Client{}
	response, err := client.Do(request)
	require.NoError(t, err)

	require.Equal(t, http.StatusNoContent, response.StatusCode)
}
