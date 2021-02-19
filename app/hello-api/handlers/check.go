package handlers

import (
	"context"
	"github.com/haibin/hello-service/foundation/web"
	"net/http"
	"os"
)

type check struct {}

func (cg check) liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	info := struct {
		Status    string `json:"status,omitempty"`
		Host      string `json:"host,omitempty"`
	}{
		Status:    "up",
		Host:      host,
	}

	return web.Respond(ctx, w, info, http.StatusOK)
}