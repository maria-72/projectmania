package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"projectmania/models"

	"github.com/gin-gonic/gin"
)

func getRequest(url string) ([]byte, error) {
	token, _ := os.LookupEnv("GITHUB_TOKEN")

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, nil
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func getConfig() models.Config {
	configJson, err := os.Open("./config.json")

	if err != nil {
		log.Println(err)
	}

	defer configJson.Close()

	var config models.Config

	byteValue, err := io.ReadAll(configJson)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteValue, &config)

	return config
}

func GetProjects(c *gin.Context) {
	url, _ := os.LookupEnv("GITHUB_API_URL")

	config := getConfig()

	resRepos, err := getRequest(fmt.Sprintf("%s/user/repos", url))

	if err != nil {
		log.Fatal(err)
	}

	var projects []models.Project

	err = json.Unmarshal(resRepos, &projects)

	if err != nil {
		log.Fatal(err)
	}

	var visibleProjects []models.Project

	for _, v := range projects {
		for _, selectedRepo := range config.Repositories {
			if selectedRepo == v.Name {
				visibleProjects = append(visibleProjects, v)
			}
		}
	}

	c.IndentedJSON(http.StatusOK, visibleProjects)
}
