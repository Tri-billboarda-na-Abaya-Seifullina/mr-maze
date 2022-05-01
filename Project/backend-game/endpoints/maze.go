package endpoints

import (
	"context"
	"github.com/Abunyawa/back_game/service"
	"github.com/go-kit/kit/endpoint"
)

type GenerateMazeRequest struct {
	Height int `json:"height"`
	Width  int `json:"width""`
}

func MakeGenerateMazeEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GenerateMazeRequest)

		return s.GenerateMaze(req.Height, req.Width)
	}
}
