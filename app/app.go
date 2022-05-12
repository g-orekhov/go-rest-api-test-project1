package app

import (
	db "g-oriekhov/testProject1/models"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type App struct {
	http   *http.Server
	Router *mux.Router
	DB     *db.DB
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
	app.DB = new(db.DB)
	if err := app.DB.ConnectDB(); err != nil {
		log.Fatal(err)
	}
	return &app
}

func GetApp() *App {
	return &app
}
