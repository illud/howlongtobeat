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
			Main:          "52h 37m",
			Extra:         "98h 12m",
			Completionist: "131h 17m",
		},
		{
			Image:         "https://howlongtobeat.com/games/108888_Elden_Ring_GB.jpg",
			Title:         "Elden Ring GB",
			Main:          "0h 21m",
			Extra:         "0h 29m",
			Completionist: "0h 0m",
		},
	}

	assert.EqualValues(t, gamesFound, expected)
}
