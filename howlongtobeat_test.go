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
			Main:          "52",
			Extra:         "98",
			Completionist: "131",
		},
		{
			Image:         "https://howlongtobeat.com/games/108888_Elden_Ring_GB.jpg",
			Title:         "Elden Ring GB",
			Main:          "21 Mins ",
			Extra:         "29 Mins ",
			Completionist: "--",
		},
	}

	assert.EqualValues(t, gamesFound, expected)
}
