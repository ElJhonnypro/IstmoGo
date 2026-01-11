package main

import (
	"API/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../internal/.env")
	routes.StartRoutes()

}
