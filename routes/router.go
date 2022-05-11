package router

import (
	"g-oriekhov/testProject1/app"
)

func RegisterRoutes(app *app.App) {
	registerEventRoutes(app)
}
