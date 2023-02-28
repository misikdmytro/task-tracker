package service

import (
	"context"

	"github.com/misikdmytro/task-tracker/internal/database"
)

type HealthService interface {
	HealthCheck(ctx context.Context) error
}

type healthService struct {
	r database.Repository
}

func NewHealthService(r database.Repository) HealthService {
	return &healthService{r: r}
}

func (h *healthService) HealthCheck(ctx context.Context) error {
	return h.r.Ping(ctx)
}
