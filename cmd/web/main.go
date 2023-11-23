package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BheeshmasenaReddy/bookings/pkg/config"
	"github.com/BheeshmasenaReddy/bookings/pkg/handlers"
	"github.com/BheeshmasenaReddy/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	app.InProd = false

	session = scs.New()
	session.Lifetime = 3 * time.Hour
	session.Cookie.Secure = app.InProd
	session.Cookie.Persist = false
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = templateCache
	app.UseCache = true

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	render.GetCache(&app)

	fmt.Println("Starting a server at port", port)

	srv := http.Server{
		Addr:    port,
		Handler: Routers(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
