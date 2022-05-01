package endpoints

import (
	"github.com/Abunyawa/back_game/service"
)

type Endpoints struct {
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{}
}
