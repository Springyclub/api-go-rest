package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Home page")
	if err != nil {
		return
	}
}

func Guigas(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Guigas")
	if err != nil {
		return
	}
}
