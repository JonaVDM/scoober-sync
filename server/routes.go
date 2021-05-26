package server

import (
	"errors"

	"github.com/jonavdm/scoober-sync/ui"
)

func (a *FactoryServer) Routes() error {
	if a.Router == nil {
		return errors.New("missing router object")
	}

	a.Router.HandleFunc("/ping", a.handlePing())
	a.Router.PathPrefix("/").Handler(ui.Serve())
	return nil
}
