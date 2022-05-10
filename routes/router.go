package router

import (
	"g-oriekhov/testProject1/app"
)

func RegisterRoutes(app *app.App) {
	registerEventRoutes(app)
	//app.Router.HandleFunc("/", indexHendler)
}

// func indexHendler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World")
// }
