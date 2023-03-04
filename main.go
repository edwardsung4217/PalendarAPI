package main

import (
	"log"
	"os"

	"github.com/VolunteerOne/volunteer-one-app/backend/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// eww
func main() {
	r := gin.Default()
	r.GET("/", handlers.RootHandler)
	r.GET("/Example", handlers.ExampleHandler)
	r.Get("/events", handlers.eventsHandler)

	r.Run()
}

// Runs before main
func init() {
	// Load Environment Variables
	err := godotenv.Load("go.env")
	if err != nil {
		log.Fatal("Could not load environment variables")
	}

	log.Printf("Application starting on port: %v", os.Getenv("PORT"))
}
