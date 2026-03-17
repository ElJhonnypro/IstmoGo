package main

import (
	"API/data"
	"API/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../internal/.env")

	data.ConnectDB()
	data.CreateTables(data.GetDB())

	routes.StartRoutes()
}
