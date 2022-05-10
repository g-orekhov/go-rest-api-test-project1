package main

import (
	"g-oriekhov/testProject1/app"
	router "g-oriekhov/testProject1/routes"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	App := app.NewApp()
	router.RegisterRoutes(App)
	log.Fatal(App.Run())
}
