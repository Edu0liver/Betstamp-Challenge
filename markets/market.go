package markets

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Market struct {
	Fixture_id string  `json:"fixture_id"`
	Bet_type   string  `json:"bet_type"`
	Is_live    bool    `json:"is_live"`
	Odds       float64 `json:"odds"`
	Number     float64 `json:"number"`
	Side_type  string  `json:"side_type"`
}

type Selection struct {
	Name string  `json:"name"`
	Odds float64 `json:"odds"`
}

type MarketEvent struct {
	MarketName string      `json:"marketName"`
	Selections []Selection `json:"selections"`
}

type Event struct {
	ID      int           `json:"id"`
	Name    string        `json:"name"`
	Start   string        `json:"start"`
	State   string        `json:"state"`
	Markets []MarketEvent `json:"markets"`
}

type APIResponse struct {
	Events []Event `json:"events"`
}

var MarketDB []Market

var marketNameMap = map[string]string{
	"Money line":    "Moneyline",
	"Points Spread": "Spread",
	"Total Points":  "Total",
}

func ProcessMarkets(apiData []byte) ([]Market, error) {
	var apiResponse APIResponse
	if err := json.Unmarshal(apiData, &apiResponse); err != nil {
		return nil, errors.New("failed to parse JSON")
	}

	var wg sync.WaitGroup
	var markets []Market
	var errors []error
	marketChan := make(chan Market, len(apiResponse.Events))
	errorChan := make(chan error)

	go processChans(&markets, &errors, marketChan, errorChan)

	for _, event := range apiResponse.Events {
		wg.Add(1)
		go processEvent(&event, marketChan, errorChan, &wg)
	}

	wg.Wait()
	close(marketChan)
	close(errorChan)

	if len(errors) > 0 {
		return markets, fmt.Errorf("encountered errors: %v", errors)
	}

	return markets, nil
}

func findFixture(team1 string, team2 string, eventDate time.Time) string {
	return fmt.Sprintf("Team 1: %s, Team 2: %s, Event Date: %v\n", team1, team2, eventDate)
}

func processChans(markets *[]Market, errors *[]error, marketChan chan Market, errorChan chan error) {
	for {
		select {

		case market := <-marketChan:
			*markets = append(*markets, market)
			MarketDB = append(MarketDB, market)

		case err := <-errorChan:
			*errors = append(*errors, err)

		}
	}
}

func processEvent(event *Event, marketCh chan<- Market, errorCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	team1, team2, found := strings.Cut(event.Name, " @ ")
	if !found {
		errorCh <- fmt.Errorf("invalid event name format: %s", event.Name)
		return
	}

	eventDate, err := time.Parse(time.RFC3339, event.Start)
	if err != nil {
		errorCh <- fmt.Errorf("invalid event date format: %s", event.Start)
		return
	}

	fixtureID := findFixture(team1, team2, eventDate)
	isLive := event.State == "LIVE"

	for _, market := range event.Markets {
		bet_type, exists := marketNameMap[market.MarketName]
		if !exists {
			errorCh <- fmt.Errorf("invalid bet_type: %s", market.MarketName)
			return
		}

		for _, selection := range market.Selections {
			number, side_type, err := getMarketType(bet_type, selection)
			if err != nil {
				errorCh <- err
				return
			}

			marketCh <- Market{
				Fixture_id: fixtureID,
				Bet_type:   bet_type,
				Is_live:    isLive,
				Odds:       selection.Odds,
				Number:     number,
				Side_type:  side_type,
			}
		}
	}

}

func getMarketType(bet_type string, selection Selection) (float64, string, error) {
	switch bet_type {

	case "Moneyline":
		return 0, selection.Name, nil

	case "Spread":
		lastSpace := strings.LastIndex(selection.Name, " ")

		if lastSpace == -1 {
			return 0, "", fmt.Errorf("invalid spread format: %s", selection.Name)
		}

		spread, err := strconv.ParseFloat(selection.Name[lastSpace+1:], 64)

		if err != nil {
			return 0, "", fmt.Errorf("invalid spread number: %s", selection.Name)
		}

		return spread, selection.Name[:lastSpace], nil

	case "Total":
		lastSpace := strings.LastIndex(selection.Name, " ")

		if lastSpace == -1 {
			return 0, "", fmt.Errorf("invalid total points format: %s", selection.Name)
		}

		total, err := strconv.ParseFloat(selection.Name[lastSpace+1:], 64)

		if err != nil {
			return 0, "", fmt.Errorf("invalid total number: %s", selection.Name)
		}

		return total, strings.ToLower(selection.Name[:lastSpace]), nil

	}

	return 0, "", nil
}
