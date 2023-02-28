package e2e_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/misikdmytro/task-tracker/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func TestCreateListLongName(t *testing.T) {
	_, start, close := Setup(t)
	defer close()
	start()

	m := model.CreateListRequest{
		Name: "SArUgjw0jT2Vpfik1ffidrsB0NopE4yplmv8YUIZmaoCPAQBViJzPmVIPVXcjuPkvIP0eB7TUE2L1uKevPAsou0zf6MMDAvZmtGKURxu9bAkbPxn399xa5heQBt11yk2F7RxVflxc6LvUR7CLZ9uGOkFtq6hgLIaaTCwvKmPt4mWKWQUaoTquMTPgzg4KtQT5HFlJndtHD9b7GCuY3WOzM9ErDFN320I72Hnq2iCj5IpuJOkuSBDUjGTSjSqNmRy1BSzbQkzTDVjYOmkfoNaKC8OSta7soPx87URGUSG5iANbyxDD2XcabEXCcETIHEMK7zAA39g0kBRuWpTfOyl67gbx4OMFvNfFo1aL2d6bAGueeDwN9ubQuHfgQEQeLtdlRtNHtgm7qYK0OKct3EsKPm51uVUfmdCzCSeOEGWBOEzXUZshBUXPS5AeGxLcpbpznhJqGrzNgM5",
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

	require.Equal(t, http.StatusBadRequest, response.StatusCode)
	var result model.ErrorResponse
	require.NoError(t, json.NewDecoder(response.Body).Decode(&result))

	assert.Equal(t, "invalid request body", result.Error)
}
