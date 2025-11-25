package main

import (
	"ticketing/backend-api/config"
	"ticketing/backend-api/database"
	"ticketing/backend-api/routes"
)

func main() {

	config.LoadEnv()
	database.InitDB()
	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
