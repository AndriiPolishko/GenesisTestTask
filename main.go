package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RateResponse struct {
	Rate int `json:"rate"`
}

func rateHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to retrieve the BTC to UAH rate from a third-party service

	var btcAPIurl string = "https://api.coingecko.com/api/v3/coins/bitcoin"
	rate, err := getBTCPriceInUAH(btcAPIurl)

	if err != nil {
		log.Fatal(err)
	}

	response := RateResponse{
		Rate: rate,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshaling JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	log.SetPrefix("Main error")

	http.HandleFunc("/api/rate", rateHandler)
	http.HandleFunc("/api/subscribe", subscribeHandler)

	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
