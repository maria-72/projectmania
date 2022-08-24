package main

import (
	"fmt"
	"log"
	"os"

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

	v1 := router.Group("/api/v1")
	{
		v1.GET("/get-projects", handlers.GetProjects)
	}

}

func main() {
	port, exists := os.LookupEnv("PORT")

	if !exists {
		port = "8081"
	}

	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}
