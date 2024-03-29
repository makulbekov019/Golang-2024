package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/real", WelcomeHandler)
	r.HandleFunc("/health", HealthCheck).Methods("GET")
	r.HandleFunc("/real/all", PrintAllPlayers).Methods("GET")
	r.HandleFunc("/real/{name}", GetPlayerByName).Methods("GET")
	http.ListenAndServe(":5500", r)
}