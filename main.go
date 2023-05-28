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

func main() {
	log.SetPrefix("Main error")

	var btcAPIurl string = "https://api.coingecko.com/api/v3/coins/bitcoin"
	priceOfBTCInUAH, err := getBTCPriceInUAH(btcAPIurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(priceOfBTCInUAH)
}
