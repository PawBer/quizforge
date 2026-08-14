package web

import (
	"net/http"

	"github.com/PawBer/quizforge"
	"github.com/go-chi/chi/v5"
)

type Application struct {
}

func (app *Application) SetupRoutes() http.Handler {
	mux := chi.NewRouter()

	mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(quizforge.StaticFS()))))

	mux.Get("/", app.GetIndex)

	return mux
}
