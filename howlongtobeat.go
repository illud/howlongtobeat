package howlongtobeat

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type GamesTitles struct {
	Title string `json:"title"`
}

type Gamesimages struct {
	Image string `json:"image"`
}

type GamesTime struct {
	Main string `json:"main"`
}

type GamesExtra struct {
	Extra string `json:"extra"`
}

type GamesCompletionist struct {
	Completionist string `json:"completionist"`
}
type Games struct {
	Image         string `json:"image"`
	Title         string `json:"title"`
	Main          string `json:"main"`
	Extra         string `json:"extra"`
	Completionist string `json:"completionist"`
}

//Filters hmtl to find games
func Search(game string) []Games {
	gamesFound := howLongToBeat(game) // gets html

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(gamesFound)))

	if err != nil {
		log.Fatal(err)
	}

	var gamesTitles []GamesTitles
	var gamesimages []Gamesimages
	var gamesTime []GamesTime
	var gamesExtra []GamesExtra
	var gamesCompletionist []GamesCompletionist

	var games []Games

	// Finds game image
	doc.Find(".back_darkish .search_list_image img").Each(func(_ int, tag *goquery.Selection) {

		link, _ := tag.Attr("src")

		gamesimages = append(gamesimages, Gamesimages{"https://howlongtobeat.com" + link})
	})

	// Finds game name
	doc.Find(".search_list_details .shadow_text a").Each(func(_ int, tag *goquery.Selection) {

		title, _ := tag.Attr("title")
		gamesTitles = append(gamesTitles, GamesTitles{title})
	})

	// Finds class .search_list_details_block and loops divs
	doc.Find(".search_list_details_block").Each(func(i int, tag *goquery.Selection) {

		// if finds .search_list_tidbit class it finds non online games else are online
		if len(tag.Find(".search_list_tidbit").Text()) > 0 {
			// If class .search_list_tidbit position is = to 1 and is not empty it finds Main
			if tag.Find(".search_list_tidbit").Eq(1).Text() == "" {
				gamesTime = append(gamesTime, GamesTime{"--"})

			} else {
				gamesTime = append(gamesTime, GamesTime{tag.Find(".search_list_tidbit").Eq(1).Text()})
			}

			if tag.Find(".search_list_tidbit").Eq(3).Text() == "" {
				gamesExtra = append(gamesExtra, GamesExtra{"--"})

			} else {
				gamesExtra = append(gamesExtra, GamesExtra{tag.Find(".search_list_tidbit").Eq(3).Text()})
			}

			if tag.Find(".search_list_tidbit").Eq(5).Text() == "" {
				gamesCompletionist = append(gamesCompletionist, GamesCompletionist{"--"})
			} else {
				gamesCompletionist = append(gamesCompletionist, GamesCompletionist{tag.Find(".search_list_tidbit").Eq(5).Text()})
			}
		} else {
			if tag.Find("div").Eq(1).Text() == "" {
				gamesTime = append(gamesTime, GamesTime{"--"})

			} else {
				gamesTime = append(gamesTime, GamesTime{tag.Find("div").Eq(1).Text()})
			}

			if tag.Find("div").Eq(3).Text() == "" {
				gamesExtra = append(gamesExtra, GamesExtra{"--"})

			} else {
				gamesExtra = append(gamesExtra, GamesExtra{tag.Find("div").Eq(3).Text()})
			}

			if tag.Find("div").Eq(5).Text() == "" {
				gamesCompletionist = append(gamesCompletionist, GamesCompletionist{"--"})

			} else {
				gamesCompletionist = append(gamesCompletionist, GamesCompletionist{tag.Find("div").Eq(5).Text()})
			}
		}

	})

	for i := range gamesTitles {
		games = append(games, Games{gamesimages[i].Image, gamesTitles[i].Title, gamesTime[i].Main, gamesExtra[i].Extra, gamesCompletionist[i].Completionist})
	}

	return games
}

func howLongToBeat(game string) string {
	form := url.Values{}
	form.Add("queryString", game)
	form.Add("t", "games")
	form.Add("sorthead", "popular")
	form.Add("sortd", "Normal Order")
	form.Add("plat", "")
	form.Add("length_type", "main")
	form.Add("length_min", "")
	form.Add("length_max", "")
	form.Add("detail", "0")
	form.Add("v", "")
	form.Add("f", "")
	form.Add("g", "")
	form.Add("randomize", "0")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest("POST", "https://howlongtobeat.com/search_results?page=1", strings.NewReader(form.Encode()))

	if err != nil {
		//handle postform error
		fmt.Println(err)
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
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
