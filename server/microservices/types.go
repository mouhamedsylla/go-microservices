package microservices

import (
	"net/http"
)

type Controller interface {
	HTTPServe() http.Handler
	EndPoint() string
}

type Host interface {
	AddController(Controller)
}
