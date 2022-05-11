package router

import (
	"g-oriekhov/testProject1/app"
	eventService "g-oriekhov/testProject1/services/events"
)

func registerEventRoutes(app *app.App) {
	eventRouter := app.Router.PathPrefix("/event/").Subrouter()
	eventRouter.Methods("GET").Path("/").HandlerFunc(eventService.GetEvents)
	eventRouter.Methods("POST").Path("/").HandlerFunc(eventService.CreateEvent)
	eventRouter.Methods("GET").Path("/{id:[0-9]+}/").HandlerFunc(eventService.GetEvent)
	eventRouter.Methods("DELETE").Path("/{id:[0-9]+}/").HandlerFunc(eventService.DeleteEvent)
	eventRouter.Methods("PUT").Path("/{id:[0-9]+}/").HandlerFunc(eventService.UpdateEvent)
}
