package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BTCPrice struct {
	MarketData MarketData `json:"market_data"`
}

type MarketData struct {
	CurrentPrice CurrentPrice `json:"current_price"`
}

type CurrentPrice struct {
	UAH int `json:"uah"`
}

func getBTCPriceInUAH(url string) (int, error) {
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	var btcPrice BTCPrice
	err = json.NewDecoder(response.Body).Decode(&btcPrice)
	if err != nil {
		return 0, err
	}

	return btcPrice.MarketData.CurrentPrice.UAH, nil
}

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

	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
