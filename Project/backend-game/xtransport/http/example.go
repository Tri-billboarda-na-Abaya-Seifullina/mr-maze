package xhttp

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Abunyawa/back_game/endpoints"
)

func decodeHTTPExampleRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := &endpoints.ExampleRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, err
	}
	return req, nil
}
