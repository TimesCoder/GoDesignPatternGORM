package main

import (
	"design_pattern/pkg/config"
	"design_pattern/pkg/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize the database
	config.InitDB()

	// Initialize Echo instance
	e := echo.New()

	// Initialize routes
	routes.InitRoutes(e)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
