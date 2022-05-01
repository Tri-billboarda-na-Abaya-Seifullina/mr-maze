package endpoints

import (
	"github.com/Abunyawa/back_game/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GenerateMazeEndpoint endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		GenerateMazeEndpoint: MakeGenerateMazeEndpoint(s),
	}
}
