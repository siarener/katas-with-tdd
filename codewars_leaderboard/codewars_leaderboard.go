package kata

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type LeaderboardEntry struct {
	Rank     int
	Username string
	Honor    int
}

func getWebageContent(address string) *goquery.Document {
	response, err := http.Get(address)
	if err != nil {
		log.Fatal("Error making HTML request:", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatal("Error: Unexpected HTTP status code:", response.Status)
	}

	content, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error parsing HTML: ", err)
	}

	return content
}

func GetLeaderboardHonor() []LeaderboardEntry {
	apiURL := "https://www.codewars.com/users/leaderboard"

	webpageContent := getWebageContent(apiURL)

	table := webpageContent.Find("table")

	var leaderboard []LeaderboardEntry

	table.Find("tr").Each(func(rowIndex int, rowHtml *goquery.Selection) {

		var entry LeaderboardEntry

		rank := strings.TrimSpace(rowHtml.Find("td.rank").Text())

		if rank != "" {
			// get rank
			rankInt, err := strconv.Atoi(strings.TrimPrefix(rank, "#"))
			if err != nil {
				log.Fatal("Error parsing rank information: ", err)
			}
			entry.Rank = rankInt

			// get username
			entry.Username = rowHtml.AttrOr("data-username", "")

			// get honor amount
			honorStr := strings.TrimSpace(rowHtml.Find("td.honor").Text())
			honorInt, err := strconv.Atoi(strings.ReplaceAll(honorStr, ",", ""))
			if err != nil {
				log.Fatal("Error parsing honor information: ", err)

			}
			entry.Honor = honorInt

			// add entry
			leaderboard = append(leaderboard, entry)
		}
	})

	return leaderboard
}
