package xhttp

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Abunyawa/back_game/domain"
	"github.com/Abunyawa/back_game/endpoints"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

var (
	ErrorBadRequest = errors.New("bad request")
)

func decodeHTTPGenerateMazeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := &endpoints.GenerateMazeRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.WithFields(log.Fields{
			"method": domain.GENERATING,
		}).Error("error reading request body")
		return nil, err
	}

	log.WithFields(log.Fields{
		"method": domain.GENERATING,
	}).Info("Got request")
	return req, nil
}

func decodeHTTPGetMazeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := &endpoints.GetMazeRequest{}

	idString := r.URL.Query().Get("id")

	if idString == "" {
		log.WithFields(log.Fields{
			"method": domain.READING,
		}).Error("error bad request")
		return nil, ErrorBadRequest
	}
	var id int64

	if id, err = strconv.ParseInt(idString, 10, 32); err != nil {
		log.WithFields(log.Fields{
			"method": domain.READING,
		}).Error("error bad request")
		return nil, ErrorBadRequest
	}

	req.Id = int(id)
	log.WithFields(log.Fields{
		"method": domain.READING,
	}).Info("Got request")
	return req, nil
}
