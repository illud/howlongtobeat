package howlongtobeat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Games struct {
	Image         string `json:"image"`
	Title         string `json:"title"`
	Main          string `json:"main"`
	Extra         string `json:"extra"`
	Completionist string `json:"completionist"`
}

type HowlongtobeatResponse struct {
	Data []HowlongtobeatResponseData `json:"data"`
}

type HowlongtobeatResponseData struct {
	Game_image string `json:"game_image"`
	Game_name  string `json:"game_name"`
	Comp_main  int    `json:"comp_main"`
	Comp_plus  int    `json:"comp_plus"`
	Comp_100   int    `json:"comp_100"`
}

func secondsToTime(e int) string {
	h := e / 3600
	m := e % 3600 / 60

	var hours string = strconv.Itoa(h)

	var minutes string = strconv.Itoa(m)

	return hours + "h " + minutes + "m"
}

//Filters hmtl to find games
func Search(game string) []Games {
	gamesFound := howLongToBeat(game) // gets html

	var howlongtobeatResponse HowlongtobeatResponse

	err := json.Unmarshal([]byte(gamesFound), &howlongtobeatResponse)
	if err != nil {
		fmt.Println(err)
	}

	var games []Games

	for i := range howlongtobeatResponse.Data {
		games = append(games, Games{"https://howlongtobeat.com/games/" + howlongtobeatResponse.Data[i].Game_image, howlongtobeatResponse.Data[i].Game_name, secondsToTime(howlongtobeatResponse.Data[i].Comp_main), secondsToTime(howlongtobeatResponse.Data[i].Comp_plus), secondsToTime(howlongtobeatResponse.Data[i].Comp_100)})
	}

	return games
}

func howLongToBeat(game string) string {
	var jsonData = []byte(`{
		"searchType": "games",
		"searchTerms": [
		    "` + game + `"
		],
		"searchPage": 1,
		"size": 20,
		"searchOptions": {
		  "games": {
			"userId": 0,
			"platform": "",
			"sortCategory": "popular",
			"rangeCategory": "main",
			"rangeTime": {
			  "min": 0,
			  "max": 0
			},
			"gameplay": {
			  "perspective": "",
			  "flow": "",
			  "genre": ""
			},
			"modifier": ""
		  },
		  "users": {
			"sortCategory": "postcount"
		  },
		  "filter": "",
		  "sort": 0,
		  "randomizer": 0
		}
	  }`)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest("POST", "https://www.howlongtobeat.com/api/search", bytes.NewBuffer(jsonData))

	if err != nil {
		//handle postform error
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("origin", "https://howlongtobeat.com")
	req.Header.Add("referer", "https://howlongtobeat.com")

	resp, err := client.Do(req)

	if err != nil {
		//handle postform error
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		//handle read response error
		fmt.Println(err)
	}

	return string(body)
}
