package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// /wallets - GET, POST
// /wallets/{ID} - GET, PUT, DELETE

// /exchange - GET

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/wallets", getWallets).
		Methods(http.MethodGet)

	log.Default().Println("Starting server ...")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func getWallets(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("Handling HTTP request")

	wallets := []Wallet{
		{
			ID:         0,
			HolderName: "John",
			Amount:     10,
		},
		{
			ID:         1,
			HolderName: "Peter",
			Amount:     0,
		},
	}

	err := json.NewEncoder(w).Encode(wallets)
	if err != nil {
		log.Default().Printf("Something went wrong while writing response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
