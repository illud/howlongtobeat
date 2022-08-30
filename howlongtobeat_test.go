package howlongtobeat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {

	gamesFound := Search("Elden Ring")

	if len(gamesFound) == 0 {
		t.Error("No games Found")
	}

	type Games struct {
		Image         string `json:"image"`
		Title         string `json:"title"`
		Main          string `json:"main"`
		Extra         string `json:"extra"`
		Completionist string `json:"completionist"`
	}

	expected := []Games{
		{
			Image:         "https://howlongtobeat.com/games/68151_Elden_Ring.jpg",
			Title:         "Elden Ring",
			Main:          "Main Story 52 Hours ",
			Extra:         "Main + Extra 98 Hours ",
			Completionist: "Completionist 131 Hours ",
		},
		{
			Image:         "https://howlongtobeat.com/games/108888_Elden_Ring_GB.jpg",
			Title:         "Elden Ring GB",
			Main:          "Main Story 21 Mins ",
			Extra:         "Main + Extra 29 Mins ",
			Completionist: "Completionist --",
		},
	}

	assert.EqualValues(t, gamesFound, expected)
}
