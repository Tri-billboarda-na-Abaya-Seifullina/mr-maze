package xhttp

import (
	"context"
	"encoding/json"
	"github.com/Abunyawa/back_game/domain"
	"github.com/Abunyawa/back_game/endpoints"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func decodeHTTPGenerateMaze(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := &endpoints.AddUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.WithFields(log.Fields{
			"method": domain.REGISTER,
		}).Error("error reading request body")
		return nil, err
	}

	log.WithFields(log.Fields{
		"method": domain.REGISTER,
		"login":  req.Login,
	}).Info("Got request")
	return req, nil
}
