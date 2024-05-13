package microservices

import (
	"realtimeForum/server/router"
)

type Service interface {
	GetService() *Microservice
	InitService()
	ConfigureEndpoint()
}

type AppServices struct {
	Microservices []Service
}

func (aps *AppServices) InitServices() {
	for _, service := range aps.Microservices {
		service.InitService()
		service.ConfigureEndpoint()
	}
}

func NewAppServices(services ...Service) *AppServices {
	aps := &AppServices{}
	aps.Microservices = append(aps.Microservices, services...)
	return aps
}

type Microservice struct {
	ServiceName string
	Router      *router.Router
	Controllers []Controller
	Port        string
}

func NewMicroservice(name, port string) *Microservice {
	return &Microservice{
		ServiceName: name,
		Port:        port,
		Router:      router.NewRouter(),
	}
}

func (m *Microservice) AddController(c Controller) {
	m.Controllers = append(m.Controllers, c)
}
