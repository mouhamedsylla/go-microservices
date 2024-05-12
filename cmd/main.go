package main

import (
	"realtimeForum/auth"
	realtimechat "realtimeForum/realtimeChat"
	"realtimeForum/server"
	"realtimeForum/server/microservices"
)

func main() {
	allServices := []microservices.Service{
		&auth.Auth{},
		&realtimechat.Chat{},
	}
	server := server.NewServer(allServices...)
	server.StartServices("8080")
}
