package xhttp

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"

	"github.com/Abunyawa/back_game/endpoints"
	"github.com/Abunyawa/back_game/service"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(s service.Service) http.Handler {

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}

	r := mux.NewRouter()
	e := endpoints.MakeEndpoints(s)

	example := httptransport.NewServer(
		e.ExampleEndpoint,
		decodeHTTPExampleRequest,
		encodeResponse,
		options...,
	)

	addUser := httptransport.NewServer(
		e.AddUserEndpoint,
		decodeHTTPAddUserRequest,
		encodeResponse,
		options...,
	)

	authUser := httptransport.NewServer(
		e.AuthUserEndpoint,
		decodeHTTPAuthUserRequest,
		encodeResponse,
		options...,
	)

	refreshToken := httptransport.NewServer(
		e.RefreshTokenEndpoint,
		decodeHTTPRefreshTokenRequest,
		encodeResponse,
		options...,
	)

	r.Handle("/example", example).Methods("POST")
	r.Handle("/register", addUser).Methods("POST")
	r.Handle("/auth", authUser).Methods("POST")
	r.Handle("/refresh", refreshToken).Methods("POST")

	return r
}

type errorWrapper struct {
	Error string `json:"error"`
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func err2code(err error) int {
	switch err {
	case service.ErrorUnauthorized:
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
