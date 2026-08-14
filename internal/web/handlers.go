package web

import (
	"context"
	"net/http"

	"github.com/PawBer/quizforge/templates"
)

func (app *Application) GetIndex(w http.ResponseWriter, req *http.Request) {
	err := templates.Index().Render(context.TODO(), w)
	if err != nil {
		return
	}
}
