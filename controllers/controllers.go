package controllers

import (
	"api-rest/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Home page")
	if err != nil {
		return
	}
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(models.Personalities)
	if err != nil {
		return
	}
}

func ReturnPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	for _, personalities := range models.Personalities {
		if strconv.Itoa(personalities.Id) == id {
			err := json.NewEncoder(w).Encode(personalities)
			if err != nil {
				return
			}
		}
	}

}
