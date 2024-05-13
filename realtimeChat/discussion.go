package realtimechat

import (
	"fmt"
	"net/http"
)

func (d *Discussion) HTTPServe() http.Handler {
	return http.HandlerFunc(d.Discussion)
}

func (d *Discussion) EndPoint() string {
	return "/discussion"
}

func (d *Discussion) Discussion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the discussion microservices...")
}
