package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"server/database"
	"server/model"

	"github.com/joho/godotenv"
)

type test_struct struct {
	Test string
}

func testEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	myName := os.Getenv("APP_NAME")
	log.Println(myName)
}

func main() {
	testEnv()
	database.TestDatabaseService()
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/add_budget", addBudgetHandler)
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

// statusHandler : Example to make other handlers
func statusHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}

// addBudgetHandler: add a budget and calculate installments
func addBudgetHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var financing model.Budget
	err = json.Unmarshal(b, &financing)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	financing = model.CalcInstallments(financing)
	database.SaveBudget(financing)
	//concatenated := fmt.Sprintf("{installment_qty: %s, installment_value: %b}", financing.Installment, financing.InstallmentValue)
	var buf []byte
	buf, err = json.Marshal(financing)
	w.Write(buf)
}
