package endpoints

import (
	"context"

	"github.com/Abunyawa/back_game/service"
	"github.com/go-kit/kit/endpoint"
)

type ExampleRequest struct {
	Name string `json:"name"`
}

type ExampleResponse struct {
	Phrase string `json:"phrase"`
}

func MakeExampleEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*ExampleRequest)

		res, err := s.ExampleServiceMethod(req.Name)

		if err != nil {
			return nil, err
		}

		return ExampleResponse{
			Phrase: res,
		}, nil
	}
}
