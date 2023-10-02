package main

import (
	_ "github.com/gin-gonic/gin"
	_ "github.com/labstack/echo/v4"
	_ "gopkg.in/reform.v1"
	"local/api"
	"local/database"
	"log"
)

// main is the entry point of the program.
// It initializes the database, sets up the router and starts the server.
func main() {
	// Initialize the database
	err := database.InitDatabase()
	if err != nil {
		// Log the error and exit
		log.Fatalln("could not create database", err)
	}

	// Set up the router
	r := api.SetupRouter()
	r.Logger.Fatal(r.Start(":1323"))
}
