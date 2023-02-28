package main

import (
	"fmt"
	"os"

	"github.com/misikdmytro/task-tracker/internal/bootstrap"
)

//	@title			Task Tracker API
//	@version		1.0
//	@description	Task tracker service.
//	@contact.name	Misik Dmytro
//	@contact.url	https://github.com/misikdmytro
//	@BasePath		/
func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	s, err := bootstrap.NewServer(fmt.Sprintf("./config/%s.config.yaml", env))
	if err != nil {
		panic(err)
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
