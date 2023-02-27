package model

import "time"

type TaskListDto struct {
	ListID        int       `db:"list_id"`
	ListName      string    `db:"list_name"`
	ListCreatedAt time.Time `db:"list_created_at"`

	TaskID        *int       `db:"task_id"`
	TaskName      *string    `db:"task_name"`
	TaskCreatedAt *time.Time `db:"task_created_at"`
}
