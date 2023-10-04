package handler

import (
	"net/http"

	"github.com/Zdev147/room_booking/pkg/config"
	"github.com/Zdev147/room_booking/pkg/models"
	"github.com/Zdev147/room_booking/pkg/render"
)

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// Repository
var Repo *Repository

type Repository struct {
	appConfig *config.AppConfig
}

func CreateNewRepo(a *config.AppConfig) Repository {
	return Repository{
		appConfig: a,
	}
}

func CreateNewHanlder(r *Repository) {
	Repo = r
}
