package main

import (
	"api-rest/models"
	"api-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{
			Name:    "Guilherme Fernandes dos Santos",
			History: "História da minha vida",
			Id:      2,
		},
		{
			Name:    "Guigas",
			History: "História",
			Id:      1,
		},
	}

	routes.HandleRequest()

}
