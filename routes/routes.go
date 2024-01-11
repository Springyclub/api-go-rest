package routes

import (
	"api-rest/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/guigas", controllers.Guigas)

	log.Fatal(http.ListenAndServe(":8000", nil))
}