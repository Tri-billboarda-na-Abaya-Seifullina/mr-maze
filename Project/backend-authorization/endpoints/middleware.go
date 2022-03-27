package endpoints

import (
	"github.com/Abunyawa/back_auth/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ExampleEndpoint endpoint.Endpoint
	AddUserEndpoint endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		ExampleEndpoint: MakeExampleEndpoint(s),
		AddUserEndpoint: MakeAddUserEndpoint(s),
	}
}
