package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Application struct {
}

func (app *Application) SetupRoutes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", app.GetIndex)

	return mux
}
