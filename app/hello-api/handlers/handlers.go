package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/haibin/hello-service/business/mid"
	"github.com/haibin/hello-service/foundation/web"
)

func API(shutdown chan os.Signal, log *log.Logger) *web.App {
	// logger comes before errors
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log))

	check := check{}
	// We want check.liveness to return errors.
	app.Handle(http.MethodGet, "/liveness", check.liveness)

	return app
}
