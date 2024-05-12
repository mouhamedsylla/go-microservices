package utils

import (
	"net/http"
)

type Controller interface {
	HTTPServer() http.Handler
	EndPoint() string
}

type Host interface {
	AddController(Controller)
}
