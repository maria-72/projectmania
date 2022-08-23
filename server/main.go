package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"projectmania/handlers"
)

var router *gin.Engine

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	router = gin.Default()

	router.GET("/get-projects", handlers.GetProjects)
}

func main() {
	log.Fatal(router.Run("localhost:8080"))
}
