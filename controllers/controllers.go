package controllers

import (
	"api-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
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
