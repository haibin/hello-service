package mid

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/haibin/hello-service/foundation/web"
	//"time"
)

func Logger(log *log.Logger) web.Middleware {
	return func(handler web.Handler) web.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// If the context is missing this value, request the service
			// to be shutdown gracefully.
			v, ok := ctx.Value(web.KeyValues).(*web.Values)
			if !ok {
				return web.NewShutdownError("web value missing from context")
			}

			log.Printf("%s: started: %s %s -> %s",
				v.TraceID,
				r.Method, r.URL.Path, r.RemoteAddr,
			)

			// Call the next handler.
			err := handler(ctx, w, r)

			log.Printf("%s: completed: %s %s -> %s (%d) (%s)",
				v.TraceID,
				r.Method, r.URL.Path, r.RemoteAddr,
				v.StatusCode, time.Since(v.Now),
			)

			// Return the error so it can be handled further up the chain.
			return err
		}
	}
}
