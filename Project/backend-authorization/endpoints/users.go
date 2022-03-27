package endpoints

import (
	"context"

	"github.com/Abunyawa/back_auth/domain"
	"github.com/Abunyawa/back_auth/service"
	"github.com/go-kit/kit/endpoint"
)

type AddUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AddUserResponse struct {
	Phrase string `json:"phrase"`
}

func MakeAddUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*AddUserRequest)

		err = s.AddUser(&domain.User{
			Login:    req.Login,
			Password: req.Password,
		})

		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}
