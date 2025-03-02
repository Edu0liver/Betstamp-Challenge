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
		log.Fatalf("Error reading file: %v\n\n", err)
	}

	marketsProcessed, err := markets.ProcessMarkets(jsonData)

	fmt.Printf("\nNumber of markets processed: %+v\n\n", len(marketsProcessed))

	if err != nil {
		fmt.Printf("Error processing markets: %v", err)
	}

	for _, market := range marketsProcessed {
		fmt.Printf("%+v\n\n", market)
	}
}
