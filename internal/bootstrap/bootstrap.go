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
	F database.ConnectionFactory
	R database.Repository
	S service.ListService
}

func NewServer(in string) (*Server, error) {
	c, err := config.NewConfig(in)
	if err != nil {
		return nil, err
	}

	f := database.NewConnectionFactory(c.Database)
	r := database.NewRepository(f)
	s := service.NewListService(r)
	h := handler.NewListHandler(s)
	srvr := server.NewServer(h)

	return &Server{
		Server: srvr,
		F:      f,
		R:      r,
		S:      s,
	}, nil
}
