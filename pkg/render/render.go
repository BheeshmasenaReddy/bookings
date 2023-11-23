package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/BheeshmasenaReddy/bookings/pkg/config"
	"github.com/BheeshmasenaReddy/bookings/pkg/models"
)

func GetTemplateData(td *models.TemplateData) *models.TemplateData {
	return td
}

var app *config.AppConfig

func GetCache(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// Create template cache
	var templateCache map[string]*template.Template
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}
	// get template cache
	temp, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not find remplate")
	}

	buf := new(bytes.Buffer)

	td = GetTemplateData(td)
	err := temp.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}

	// load template cache
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	templates, err := filepath.Glob("./templates/*page.tmpl")

	if err != nil {
		return cache, err
	}

	for _, page := range templates {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return cache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			//ts, err = ts.ParseGlob("./templates/*layout.tmpls")
			ts, err = ts.ParseFiles(layouts...)

			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts

	}
	return cache, nil

}
