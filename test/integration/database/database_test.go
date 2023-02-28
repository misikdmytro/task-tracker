package database_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/misikdmytro/task-tracker/internal/database"
	"github.com/misikdmytro/task-tracker/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateList(t *testing.T) {
	c := RequireConfig(t)
	f := database.NewConnectionFactory(c.Database)
	r := database.NewRepository(f)

	name := uuid.NewString()
	id, err := r.CreateList(context.Background(), name)
	require.NoError(t, err)

	assert.Greater(t, id, 0)

	db, err := f.NewDB()
	require.NoError(t, err)
	defer db.Close()

	var result string
	require.NoError(t, db.Get(&result, "SELECT name FROM tbl_lists WHERE id = $1", id))
	assert.Equal(t, name, result)
}

func TestGetListNoList(t *testing.T) {
	c := RequireConfig(t)
	f := database.NewConnectionFactory(c.Database)

	r := database.NewRepository(f)
	result, err := r.GetTaskList(context.Background(), -1)
	require.NoError(t, err)

	assert.Equal(t, []model.TaskListDto{}, result)
}

func TestGetListEmptyList(t *testing.T) {
	c := RequireConfig(t)
	f := database.NewConnectionFactory(c.Database)

	db, err := f.NewDB()
	require.NoError(t, err)
	defer db.Close()

	name := uuid.NewString()
	var id int
	require.NoError(t, db.Get(&id, "INSERT INTO tbl_lists (name) VALUES ($1) RETURNING id", name))

	var createdAt time.Time
	require.NoError(t, db.Get(&createdAt, "SELECT created_at FROM tbl_lists WHERE id = $1", id))

	r := database.NewRepository(f)
	result, err := r.GetTaskList(context.Background(), id)
	require.NoError(t, err)

	assert.Equal(t, []model.TaskListDto{{ListID: id, ListName: name, ListCreatedAt: createdAt}}, result)
}

func TestGetList(t *testing.T) {
	c := RequireConfig(t)
	f := database.NewConnectionFactory(c.Database)

	db, err := f.NewDB()
	require.NoError(t, err)
	defer db.Close()

	name := uuid.NewString()
	var id int
	require.NoError(t, db.Get(&id, "INSERT INTO tbl_lists (name) VALUES ($1) RETURNING id", name))

	var createdAt time.Time
	require.NoError(t, db.Get(&createdAt, "SELECT created_at FROM tbl_lists WHERE id = $1", id))

	taskName := uuid.NewString()
	var taskId int
	require.NoError(t, db.Get(&taskId, "INSERT INTO tbl_tasks (name, list_id) VALUES ($1, $2) RETURNING id", taskName, id))

	var taskCreatedAt time.Time
	require.NoError(t, db.Get(&taskCreatedAt, "SELECT created_at FROM tbl_tasks WHERE id = $1", taskId))

	r := database.NewRepository(f)
	result, err := r.GetTaskList(context.Background(), id)
	require.NoError(t, err)

	assert.Equal(t, []model.TaskListDto{
		{ListID: id, ListName: name, ListCreatedAt: createdAt, TaskID: &taskId, TaskName: &taskName, TaskCreatedAt: &taskCreatedAt},
	}, result)
}

func TestCreateTask(t *testing.T) {
	c := RequireConfig(t)
	f := database.NewConnectionFactory(c.Database)

	db, err := f.NewDB()
	require.NoError(t, err)
	defer db.Close()

	listName := uuid.NewString()
	var listID int
	require.NoError(t, db.Get(&listID, "INSERT INTO tbl_lists (name) VALUES ($1) RETURNING id", listName))

	r := database.NewRepository(f)

	name := uuid.NewString()
	id, err := r.CreateTask(context.Background(), listID, name)
	require.NoError(t, err)

	assert.Greater(t, id, 0)

	var result string
	require.NoError(t, db.Get(&result, "SELECT name FROM tbl_tasks WHERE id = $1", id))
	assert.Equal(t, name, result)
}

func TestCreateTaskNoList(t *testing.T) {
	c := RequireConfig(t)
	f := database.NewConnectionFactory(c.Database)

	r := database.NewRepository(f)

	name := uuid.NewString()
	id, err := r.CreateTask(context.Background(), -1, name)
	require.Error(t, err)
	assert.Equal(t, id, 0)
}

func TestDeleteTask(t *testing.T) {
	c := RequireConfig(t)
	f := database.NewConnectionFactory(c.Database)

	db, err := f.NewDB()
	require.NoError(t, err)
	defer db.Close()

	listName := uuid.NewString()
	var listID int
	require.NoError(t, db.Get(&listID, "INSERT INTO tbl_lists (name) VALUES ($1) RETURNING id", listName))

	name := uuid.NewString()
	var id int
	require.NoError(t, db.Get(&id, "INSERT INTO tbl_tasks (name, list_id) VALUES ($1, $2) RETURNING id", name, listID))

	r := database.NewRepository(f)

	err = r.DeleteTask(context.Background(), id)
	require.NoError(t, err)

	var result int
	require.NoError(t, db.Get(&result, "SELECT COUNT(*) FROM tbl_tasks WHERE id = $1", id))
	assert.Equal(t, 0, result)
}
