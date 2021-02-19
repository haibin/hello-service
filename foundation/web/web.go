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
	//mw []Middleware
}

func NewApp(shutdown chan os.Signal) *App {
	return &App{
		ContextMux: httptreemux.NewContextMux(),
		shutdown: shutdown,
	}
}

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func (a *App) Handle(method string, path string, handler Handler) {
	h := func(w http.ResponseWriter, r *http.Request){
		ctx := r.Context()

		if err := handler(ctx, w, r); err != nil {
			//a.SignalShutdown();
			return
		}
	}
	a.ContextMux.Handle(method, path, h)
}