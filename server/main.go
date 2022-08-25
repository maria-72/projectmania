package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
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
}

func useMiddleware() {
	origins, originsExists := os.LookupEnv("ORIGINS")
	var originsSlice []string

	if !originsExists {
		originsSlice = []string{"*"}
	} else {
		originsSlice = strings.Split(origins, ",")
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     originsSlice,
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func setRoutes() {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/get-projects", handlers.GetProjects)
	}
}

func loadPort() string {
	port, portExists := os.LookupEnv("PORT")

	if !portExists {
		port = "8081"
	}

	return port
}

func main() {
	port := loadPort()

	useMiddleware()
	setRoutes()

	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}
