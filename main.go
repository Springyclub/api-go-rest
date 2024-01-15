package main

import (
	"api-rest/models"
	"api-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{
			Name:    "Guigas",
			History: "História",
		},
		{
			Name:    "Guigas",
			History: "História",
		},
	}

	routes.HandleRequest()

}
