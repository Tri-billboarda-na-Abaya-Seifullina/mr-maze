package endpoints

import (
	"github.com/Abunyawa/back_game/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ExampleEndpoint      endpoint.Endpoint
	AddUserEndpoint      endpoint.Endpoint
	AuthUserEndpoint     endpoint.Endpoint
	RefreshTokenEndpoint endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		ExampleEndpoint:      MakeExampleEndpoint(s),
		AddUserEndpoint:      MakeAddUserEndpoint(s),
		AuthUserEndpoint:     MakeAuthUserEndpoint(s),
		RefreshTokenEndpoint: MakeRefreshTokenEndpoint(s),
	}
}
