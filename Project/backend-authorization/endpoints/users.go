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
	Result string `json:"result"`
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

		return AddUserResponse{
			Result: "Registered successfully",
		}, nil
	}
}

type AuthUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthUserResponse struct {
	Token string `json:"token"`
}

func MakeAuthUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*AuthUserRequest)

		token, err := s.AuthUser(&domain.User{
			Login:    req.Login,
			Password: req.Password,
		})

		if err != nil {
			return nil, err
		}

		return AuthUserResponse{
			Token: token,
		}, nil
	}
}

type RefreshTokenRequest struct {
	Token string `json:"token"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}

func MakeRefreshTokenEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*RefreshTokenRequest)

		token, err := s.RefreshToken(&domain.Token{
			Token: req.Token,
		})

		if err != nil {
			return nil, err
		}

		return RefreshTokenResponse{
			Token: token,
		}, nil
	}
}
