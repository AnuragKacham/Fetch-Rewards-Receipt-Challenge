package main

import (
	"log"
	"net/http"

	"receipt-processor/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", handlers.ProcessReceipts).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")
	router.HandleFunc("/data", handlers.SetCustomData).Methods("POST")
	router.HandleFunc("/data/{key}", handlers.GetCustomData).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
