package auth

import (
	"fmt"
	"net/http"
	"realtimeForum/server/microservices"
)

type Auth struct {
	Auth *microservices.Microservice
}

func (auth *Auth) ConfigureEndpoint() {
	for _, controller := range auth.Auth.Controllers {
		fmt.Println("endpoint: ", controller.EndPoint())
		auth.Auth.Router.Method(http.MethodGet).Handler(controller.EndPoint(), controller.HTTPServe())
	}
}

func (auth *Auth) InitService() {
	controllers := []microservices.Controller{&Register{}}
	auth.Auth = microservices.NewMicroservice("Authentication", ":8080")

	for _, c := range controllers {
		auth.Auth.AddController(c)
	}
}

func (auth *Auth) GetService() *microservices.Microservice {
	return auth.Auth
}
