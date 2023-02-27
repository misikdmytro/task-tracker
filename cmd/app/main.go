package main

import "github.com/misikdmytro/task-tracker/internal/bootstrap"

func main() {
	s, err := bootstrap.NewServer("./config/config.yaml")
	if err != nil {
		panic(err)
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
