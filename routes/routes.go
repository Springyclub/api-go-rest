package routes

import (
	"api-rest/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/personalities/{id}", controllers.ReturnPersonality).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))
}
