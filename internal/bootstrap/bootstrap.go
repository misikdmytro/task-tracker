package bootstrap

import (
	"net/http"

	"github.com/misikdmytro/task-tracker/internal/config"
	"github.com/misikdmytro/task-tracker/internal/database"
	"github.com/misikdmytro/task-tracker/internal/handler"
	"github.com/misikdmytro/task-tracker/internal/server"
	"github.com/misikdmytro/task-tracker/internal/service"
)

type Server struct {
	*http.Server
	F  database.ConnectionFactory
	R  database.Repository
	LS service.ListService
	HH service.HealthService
}

func NewServer(in string) (*Server, error) {
	c, err := config.NewConfig(in)
	if err != nil {
		return nil, err
	}

	f := database.NewConnectionFactory(c.Database)
	r := database.NewRepository(f)
	l := service.NewListService(r)
	h := service.NewHealthService(r)
	lh := handler.NewListHandler(l)
	hh := handler.NewHealthHandler(h)
	srvr := server.NewServer(lh, hh)

	return &Server{
		Server: srvr,
		F:      f,
		R:      r,
		LS:     l,
		HH:     h,
	}, nil
}
