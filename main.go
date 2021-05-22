package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/jonavdm/scoober-sync/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	var s server.Server = &server.FactoryServer{
		Router: mux.NewRouter(),
	}
	if err := s.Routes(); err != nil {
		return err
	}

	return s.Serve(":3000")
}
