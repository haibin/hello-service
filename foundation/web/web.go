package web

import (
	"context"
	"github.com/dimfeld/httptreemux/v5"
	"net/http"
	"os"
)

type App struct {
	*httptreemux.ContextMux
	shutdown chan os.Signal
	mw       []Middleware
}

func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {
	return &App{
		ContextMux: httptreemux.NewContextMux(),
		shutdown:   shutdown,
		mw:         mw,
	}
}

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func (a *App) Handle(method string, path string, handler Handler) {
	// Add the application's general middleware to the handler chain.
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if err := handler(ctx, w, r); err != nil {
			//a.SignalShutdown();
			return
		}
	}
	a.ContextMux.Handle(method, path, h)
}
