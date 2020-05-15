package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"../database"
)

func main() {
	database.TestConnection()
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/add_budget", addBudgetHandler)
	http.HandleFunc("/search_budget", handleBudgetSearch)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Service Started on Port " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// enableCors : enable cors
func enableCors(w *http.ResponseWriter) {
	origin := os.Getenv("APP_ORIGIN_URL")
	if origin == "" {
		origin = "*"
	}
	(*w).Header().Set("Access-Control-Allow-Origin", origin)
}

// handleBudgetSearch : search
func handleBudgetSearch(w http.ResponseWriter, r *http.Request) {
	database.FindBudget()
}

// addBudgetHandler: a budget handler
func addBudgetHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	budget := database.Budget{104.444, 35}
	database.SaveBudget(budget)
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"add":"ok"}`)
}

// statusHandler : Example to make other handlers
func statusHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}
