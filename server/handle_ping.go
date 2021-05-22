package server

import (
	"fmt"
	"net/http"
)

func (s *FactoryServer) handlePing() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Pong")
	}
}
