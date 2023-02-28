package main

import "github.com/misikdmytro/task-tracker/internal/bootstrap"

//	@title			Task Tracker API
//	@version		1.0
//	@description	Task tracker service.
//	@contact.name	Misik Dmytro
//	@contact.url	https://github.com/misikdmytro
//	@BasePath		/
func main() {
	s, err := bootstrap.NewServer("./config/config.yaml")
	if err != nil {
		panic(err)
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
