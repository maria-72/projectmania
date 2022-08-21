package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func getRequest(url string) (string, error) {
	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	resultStringPointer := new(strings.Builder)

	_, copyErr := io.Copy(resultStringPointer, response.Body)

	if copyErr != nil {
		return "", copyErr
	}

	return resultStringPointer.String(), nil
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	url, exists := os.LookupEnv("URL")

	if exists {
		fmt.Println(url)
	}

	stringResponse, err := getRequest(url)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(stringResponse)
}
