package xhttp

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Abunyawa/back_auth/endpoints"
	"github.com/Abunyawa/back_auth/service"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(s service.Service) http.Handler {
	r := mux.NewRouter()
	e := endpoints.MakeEndpoints(s)

	example := httptransport.NewServer(
		e.ExampleEndpoint,
		decodeHTTPExampleRequest,
		encodeResponse,
	)

	addUser := httptransport.NewServer(
		e.AddUserEndpoint,
		decodeHTTPAddUserRequest,
		encodeResponse,
	)

	authUser := httptransport.NewServer(
		e.AuthUserEndpoint,
		decodeHTTPAuthUserRequest,
		encodeResponse,
	)

	r.Handle("/example", example).Methods("POST")
	r.Handle("/register", addUser).Methods("POST")
	r.Handle("/auth", authUser).Methods("POST")

	return r
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case service.ErrorUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
