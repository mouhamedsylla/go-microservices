package main

import (
	"realtimeForum/auth"
	"realtimeForum/server"
	"realtimeForum/server/microservices"
)

func main() {
	allServices := []microservices.Service{
		&auth.Auth{},
	}
	server := server.NewServer(allServices...)
	server.StartServices("8080")
}
