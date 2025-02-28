package markets

import (
	"fmt"
	"time"
)

type Market struct {
	Fixture_id string
	Bet_type   string
	Is_live    bool
	Odds       float64
	Number     float64
	Side_type  string
}

func ProcessMarkets(apiData []byte) ([]Market, error) {
	result := findFixture("Lakers", "Hornets", time.Now())

	fmt.Printf("%s", result)

	return []Market{}, nil
}

func findFixture(team1 string, team2 string, eventDate time.Time) string {
	return fmt.Sprintf("Team 1: %s, Team 2: %s, Event Date: %v\n", team1, team2, eventDate)
}
