package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Edu0liver/Betstamp-Interview-Q/markets"
)

func main() {
	jsonData, err := os.ReadFile("external-api-response.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	marketsProcessed, err := markets.ProcessMarkets(jsonData)
	if err != nil {
		log.Fatalf("Error processing markets: %v", err)
	}

	fmt.Printf("Market DB: %+v\n\n", markets.MarketDB)

	for _, market := range marketsProcessed {
		fmt.Printf("%+v\n\n", market)
	}
}
