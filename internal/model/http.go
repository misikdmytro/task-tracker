package model

import "time"

type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type List struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Tasks     []Task    `json:"tasks"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreateListRequest struct {
	Name string `json:"name" binding:"required,max=255"`
}

type CreateListResponse struct {
	ID int `json:"id"`
}

type GetListByIDResponse struct {
	List List `json:"list"`
}