package app

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type App struct {
	http   *http.Server
	Router *mux.Router
}

func (app *App) Run() error {
	return app.http.ListenAndServe()
}

// Default application
var app App

// Return pointer to App
func NewApp() *App {
	app.Router = mux.NewRouter()
	app.http = &http.Server{
		Addr:    os.Getenv("serverURL"),
		Handler: app.Router,
	}
	return &app
}

func GetApp() *App {
	return &app
}
