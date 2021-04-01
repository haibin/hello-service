package handlers

import (
	"context"
	"net/http"

	"github.com/haibin/hello-service/business/data/user"
	"github.com/haibin/hello-service/foundation/web"
	"github.com/pkg/errors"
)

type userGroup struct {
	user user.User
}

func (ug userGroup) create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	var nu user.NewUser
	if err := web.Decode(r, &nu); err != nil {
		return errors.Wrapf(err, "unable to decode payload")
	}

	usr, err := ug.user.Create(ctx, v.TraceID, nu, v.Now)
	if err != nil {
		return errors.Wrapf(err, "User: %+v", &usr)
	}

	return web.Respond(ctx, w, usr, http.StatusCreated)
}
