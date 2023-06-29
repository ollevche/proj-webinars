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

	walletRes := &walletResource{
		storage: NewStorage(),
	}

	r.HandleFunc("/wallets", walletRes.getWallets).
		Methods(http.MethodGet)

	r.HandleFunc("/wallets", walletRes.postWallet).
		Methods(http.MethodPost)

	r.HandleFunc("/wallets/{id}", walletRes.getWalletByID).
		Methods(http.MethodGet)

	log.Default().Println("Starting server ...")

	log.Fatal(http.ListenAndServe(":8080", r))
}

type walletResource struct {
	storage *Storage
}

func (wr *walletResource) getWallets(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("Handling HTTP request")

	respondWithJSON(w, wr.storage.GetAllWallets())
}

func (wr *walletResource) postWallet(w http.ResponseWriter, r *http.Request) {
	var wallet Wallet

	err := json.NewDecoder(r.Body).Decode(&wallet)
	if err != nil {
		log.Default().Printf("Failed to parse req body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := wr.storage.GetWalletByID(wallet.ID); ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	wr.storage.InsertWallet(&wallet)
}

func (wr *walletResource) getWalletByID(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("Handling HTTP request")

	walletID := mux.Vars(r)["id"]

	wallet, ok := wr.storage.GetWalletByID(walletID)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	respondWithJSON(w, wallet)
}

func respondWithJSON(w http.ResponseWriter, body any) {
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Default().Printf("Something went wrong while writing response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
