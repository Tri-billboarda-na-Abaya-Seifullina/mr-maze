package xhttp

import (
	"context"
	"encoding/json"
	"github.com/Abunyawa/back_game/endpoints"
	"github.com/Abunyawa/back_game/service"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler(s service.Service) http.Handler {

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}

	r := mux.NewRouter()
	e := endpoints.MakeEndpoints(s)

	generateMaze := httptransport.NewServer(
		e.GenerateMazeEndpoint,
		decodeHTTPGenerateMazeRequest,
		encodeResponse,
		options...,
	)

	getMaze := httptransport.NewServer(
		e.GetMazeEndpoint,
		decodeHTTPGetMazeRequest,
		encodeResponse,
		options...,
	)

	r.Handle("/generate", generateMaze).Methods("POST")
	r.Handle("/get", getMaze).Methods("GET")

	return handlers.CORS()(r)
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
	case ErrorBadRequest:
		return http.StatusBadRequest
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
