package server

import "github.com/gorilla/mux"

type Server interface {
	Routes() error
	Serve(port string) error
}

type FactoryServer struct {
	Router *mux.Router
}
