package main

import (
	"RS-Backend/api"
	"RS-Backend/config"
	"RS-Backend/models"
)

func main() {
	// Initialize the database connection.
	config.InitDB()

	// Migrate the models to database tables.
	models.MigrateDB()

	// Set up the router and start the server on port 8080.
	router := api.SetupRouter()
	router.Run(":8080")
}
