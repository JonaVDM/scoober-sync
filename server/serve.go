package server

import "net/http"

func (a *FactoryServer) Serve(port string) error {
	return http.ListenAndServe(port, a.Router)
}
