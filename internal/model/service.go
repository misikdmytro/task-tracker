package model

import "time"

type Task struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

type List struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Tasks     []Task
}
