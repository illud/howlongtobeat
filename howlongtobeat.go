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
	Time string `json:"time"`
}

type Games struct {
	Image string `json:"image"`
	Title string `json:"title"`
	Time  string `json:"time"`
}

func ParseHTML(game string) []Games {
	gamesFound := HowLongToBeat(game)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(gamesFound)))

	if err != nil {
		log.Fatal(err)
	}

	// text := doc.Find(".text_green a, title").Text()
	var gamesTitles []GamesTitles
	var gamesimages []Gamesimages
	var gamesTime []GamesTime

	var games []Games

	doc.Find(".back_darkish .search_list_image img").Each(func(_ int, tag *goquery.Selection) {

		link, _ := tag.Attr("src")

		gamesimages = append(gamesimages, Gamesimages{"https://howlongtobeat.com" + link})
	})

	doc.Find(".search_list_details .shadow_text a").Each(func(_ int, tag *goquery.Selection) {

		title, _ := tag.Attr("title")
		gamesTitles = append(gamesTitles, GamesTitles{title})
	})

	doc.Find(".search_list_details_block .search_list_tidbit").Each(func(_ int, tag *goquery.Selection) {

		time := tag.Text()
		fmt.Println(time)
		gamesTime = append(gamesTime, GamesTime{time})
	})

	for i := 0; i < len(gamesTitles); i++ {
		games = append(games, Games{gamesimages[i].Image, gamesTitles[i].Title, gamesTime[i].Time})
	}

	// json, _ := json.MarshalIndent(games, "", " ")
	// fmt.Println(string(json))
	return games
	// doc.Find(".search_list_details").Each(func(i int, s *goquery.Selection) {
	// 	// For each item found, get the title
	// 	title := s.Find("a").Text()
	// 	fmt.Printf("game %d: %s\n", i, title)

	// 	image := s.Find(".search_list_image").Text()
	// 	fmt.Printf("image %d: %s\n", i, image)

	// 	main := s.Find(".search_list_details_block div .search_list_tidbit").Text()
	// 	fmt.Println(main)
	// })

	// re := regexp.MustCompile("\\s{2,}")
	// fmt.Println(re.ReplaceAllString(text, "\n"))

	// c.OnHTML(".search_list_details", func(e *colly.HTMLElement) {
	// 	title := e.ChildAttr(".shadow_text a", "title")
	// 	fmt.Println(title)

	// 	dat := e.ChildText(".search_list_details_block(1)")
	// 	fmt.Println(dat)
	// })
}

func HowLongToBeat(game string) string {
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
