package router

import "net/http"

type Tree struct {
	node *Route
}

type Middleware func(http.Handler) http.Handler

type Router struct {
	t *Tree
	//middlewares []Middleware
	TempRoute Route
	Static    Directory
}

type Route struct {
	Label      string
	Methods    []string
	Handle     http.Handler
	Child      map[string]*Route
	Middleware []Middleware
}

type Directory struct {
	Prefix string
	Dir    http.Dir
}
