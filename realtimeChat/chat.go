package realtimechat

import (
	"net/http"
	"realtimeForum/server/microservices"
	"realtimeForum/utils"
)

type Chat struct {
	Chat *microservices.Microservice
}

func (chat *Chat) ConfigureEndpoint() {
	for _, controller := range chat.Chat.Controllers {
		chat.Chat.Router.Method(http.MethodGet).Handler(controller.EndPoint(), controller.HTTPServer())
	}
}

func (chat *Chat) InitService() {
	controllers := []utils.Controller{&Discussion{}}
	chat.Chat = microservices.NewMicroservice("Realtime Chat", ":9090")

	for _, c := range controllers {
		chat.Chat.AddController(c)
	}
}

func (chat *Chat) GetService() *microservices.Microservice {
	return chat.Chat
}
