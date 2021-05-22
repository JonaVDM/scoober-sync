package server

import "errors"

func (a *FactoryServer) Routes() error {
	if a.Router == nil {
		return errors.New("missing router object")
	}

	a.Router.HandleFunc("/ping", a.handlePing())
	return nil
}
