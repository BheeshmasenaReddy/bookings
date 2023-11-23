package handlers

import (
	"net/http"

	"github.com/BheeshmasenaReddy/bookings/pkg/config"
	"github.com/BheeshmasenaReddy/bookings/pkg/models"
	"github.com/BheeshmasenaReddy/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Welcome to Home Page")
	remoteIp := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Welcome to About Page")
	StringMap := make(map[string]string)

	StringMap["test"] = "Hi from about"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	StringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: StringMap,
	})
}
