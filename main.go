package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Edu0liver/Betstamp-Interview-Q/markets"
)

func main() {
	// jsonData := []byte(`{"events": [...]}`)

	jsonData, err := os.ReadFile("external-api-response.json")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%s", jsonData)

	markets, err := markets.ProcessMarkets(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", markets)
}
