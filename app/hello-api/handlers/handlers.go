package handlers

import (
	"github.com/haibin/hello-service/business/mid"
	"github.com/haibin/hello-service/foundation/web"
	"log"
	"net/http"
	"os"
)

func API(shutdown chan os.Signal, log *log.Logger) *web.App {
	app := web.NewApp(shutdown, mid.Errors(log))

	check := check{}
	// We want check.liveness to return errors.
	app.Handle(http.MethodGet, "/liveness", check.liveness)

	return app
}
