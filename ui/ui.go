package ui

import (
	"net/http"

	"github.com/markbates/pkger"
)

func Serve() http.Handler {
	return http.FileServer(pkger.Dir("/ui/build/"))
}
