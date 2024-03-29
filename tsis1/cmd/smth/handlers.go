package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
    "github.com/makulbekov019/Golang-2024/tsis1/pkg/models"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Real Madrid player database!\n")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "App is healthy!\n")
}

func PrintAllPlayers(w http.ResponseWriter, r *http.Request) {
	players := models.GetAllPlayers()
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(players)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetPlayerByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	player := models.GetPlayerByName(name)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}