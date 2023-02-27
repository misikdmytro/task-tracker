package e2e_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/misikdmytro/task-tracker/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetListOK(t *testing.T) {
	s, start, close := Setup(t)
	defer close()
	start()

	name := uuid.NewString()
	id, err := s.R.CreateList(context.Background(), name)
	require.NoError(t, err)

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://localhost:4000/lists/%d", id),
		nil,
	)
	require.NoError(t, err)

	client := http.Client{}
	response, err := client.Do(request)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, response.StatusCode)
	var result model.GetListByIDResponse
	require.NoError(t, json.NewDecoder(response.Body).Decode(&result))

	assert.Equal(t, id, result.List.ID)
	assert.Equal(t, name, result.List.Name)
}

func TestGetListNotFound(t *testing.T) {
	s, start, close := Setup(t)
	defer close()
	start()

	db, err := s.F.NewDB()
	require.NoError(t, err)
	defer db.Close()

	var id int
	require.NoError(t, db.Get(&id, "SELECT MAX(id) + 1 FROM tbl_lists"))

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://localhost:4000/lists/%d", id),
		nil,
	)
	require.NoError(t, err)

	client := http.Client{}
	response, err := client.Do(request)
	require.NoError(t, err)

	require.Equal(t, http.StatusNotFound, response.StatusCode)
	var result model.ErrorResponse
	require.NoError(t, json.NewDecoder(response.Body).Decode(&result))

	assert.Equal(t, "list not found", result.Error)
}

func TestCreateList(t *testing.T) {
	s, start, close := Setup(t)
	defer close()
	start()

	m := model.CreateListRequest{
		Name: uuid.NewString(),
	}

	jsonBytes, err := json.Marshal(m)
	require.NoError(t, err)

	request, err := http.NewRequest(
		http.MethodPut,
		"http://localhost:4000/lists/",
		bytes.NewReader(jsonBytes),
	)
	require.NoError(t, err)

	client := http.Client{}
	response, err := client.Do(request)
	require.NoError(t, err)

	require.Equal(t, http.StatusCreated, response.StatusCode)
	var result model.CreateListResponse
	require.NoError(t, json.NewDecoder(response.Body).Decode(&result))

	assert.Greater(t, result.ID, 0)

	db, err := s.F.NewDB()
	require.NoError(t, err)
	defer db.Close()

	var name string
	require.NoError(t, db.Get(&name, "SELECT name FROM tbl_lists WHERE id = $1", result.ID))
	assert.Equal(t, m.Name, name)
}
