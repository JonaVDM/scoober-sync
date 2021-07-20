package server

import (
	"encoding/json"
	"io"
	"net/http"
)

func (a *FactoryServer) handleAuth() http.HandlerFunc {
	type response struct {
		Error string `json:"error,omitempty"`
		Data  string `json:"data,omitempty"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		token := string(data)

		if err != nil {
			json.NewEncoder(rw).Encode(response{
				Error: "Something went wrong",
			})

			return
		}

		json.NewEncoder(rw).Encode(response{
			Data: token,
		})
	}
}
