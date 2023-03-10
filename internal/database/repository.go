package database

import (
	"context"
	"fmt"

	"github.com/lib/pq"
	"github.com/misikdmytro/task-tracker/internal/model"
)

var (
	ErrListForeignKeyViolation = fmt.Errorf("list not found")
)

type Repository interface {
	GetTaskList(ctx context.Context, listID int) ([]model.TaskListDto, error)
	CreateList(ctx context.Context, name string) (int, error)
	CreateTask(ctx context.Context, listID int, name string) (int, error)
	DeleteTask(ctx context.Context, taskID int) error
	Ping(ctx context.Context) error
}

type repository struct {
	f ConnectionFactory
}

func NewRepository(f ConnectionFactory) Repository {
	return &repository{f: f}
}

func (r *repository) GetTaskList(ctx context.Context, listID int) ([]model.TaskListDto, error) {
	const query = `SELECT l.id AS list_id,
		l.name AS list_name,
		l.created_at AS list_created_at,
		t.id AS task_id,
		t.name AS task_name,
		t.created_at AS task_created_at
		FROM tbl_lists AS l
		LEFT JOIN tbl_tasks AS t ON t.list_id = l.id
		WHERE l.id = $1`

	db, err := r.f.NewDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result := []model.TaskListDto{}
	if err := db.SelectContext(ctx, &result, query, listID); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) CreateList(ctx context.Context, name string) (int, error) {
	const query = `INSERT INTO tbl_lists (name) VALUES ($1) RETURNING id`

	db, err := r.f.NewDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var id int
	if err := db.GetContext(ctx, &id, query, name); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) CreateTask(ctx context.Context, listID int, name string) (int, error) {
	const query = `INSERT INTO tbl_tasks (list_id, name) VALUES ($1, $2) RETURNING id`

	db, err := r.f.NewDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var id int
	if err := db.GetContext(ctx, &id, query, listID, name); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			return 0, ErrListForeignKeyViolation
		}

		return 0, err
	}

	return id, nil
}

func (r *repository) DeleteTask(ctx context.Context, taskID int) error {
	const query = `DELETE FROM tbl_tasks WHERE id = $1`

	db, err := r.f.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.ExecContext(ctx, query, taskID); err != nil {
		return err
	}

	return nil
}

func (r *repository) Ping(ctx context.Context) error {
	db, err := r.f.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	return db.PingContext(ctx)
}
