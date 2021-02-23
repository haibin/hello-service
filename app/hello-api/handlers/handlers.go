package handlers

import (
	"github.com/haibin/hello-service/business/data/user"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"

	"github.com/haibin/hello-service/business/mid"
	"github.com/haibin/hello-service/foundation/web"
)

func API(shutdown chan os.Signal, log *log.Logger, db *sqlx.DB) *web.App {
	// Logger comes before Errors
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log))

	check := check{}
	// We want check.liveness to return errors.
	app.Handle(http.MethodGet, "/liveness", check.liveness)

	// Register user management and authentication endpoints.
	ug := userGroup{
		user: user.New(log, db),
	}
	app.Handle(http.MethodPost, "/v1/users", ug.create)

	return app
}
