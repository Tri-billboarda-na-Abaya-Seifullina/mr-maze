package xhttp

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Abunyawa/back_auth/domain"
	"github.com/Abunyawa/back_auth/endpoints"
	log "github.com/sirupsen/logrus"
)

func decodeHTTPAddUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
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

func decodeHTTPAuthUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := &endpoints.AuthUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.WithFields(log.Fields{
			"method": domain.AUTH,
		}).Error("error reading request body")
		return nil, err
	}

	log.WithFields(log.Fields{
		"method": domain.AUTH,
		"login":  req.Login,
	}).Info("Got request")
	return req, nil
}

func decodeHTTPRefreshTokenRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := &endpoints.RefreshTokenRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.WithFields(log.Fields{
			"method": domain.REFRESH,
		}).Error("error reading request body")
		return nil, err
	}

	log.WithFields(log.Fields{
		"method": domain.REFRESH,
		"token":  req.Token,
	}).Info("Got request")
	return req, nil
}
