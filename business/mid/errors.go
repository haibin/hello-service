package mid

import (
	"context"
	"github.com/haibin/hello-service/foundation/web"
	"log"
	"net/http"
)

func Errors(log *log.Logger) web.Middleware {
	// This is the actual middleware function to be executed.
	return func(handler web.Handler) web.Handler {
		// Create the handler that will be attached in the middleware chain.
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// Run the next handler and catch any propagated error.
			if err := handler(ctx, w, r); err != nil {
				// Respond to the error.
				if err := web.RespondError(ctx, w, err); err != nil {
					return err
				}
			}
			// The error has been handled so we can stop propagating it.
			return nil
		}
	}
}
