package ptc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//apiVersion is a simple endpoint handler that
//just writes a simple struct as json as the HTTP response.
func apiVersion(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		}{"PTC-Search-Service", "0.1"})
}

func getHastags(w http.ResponseWriter, r *http.Request) {

	var limit int = 100
	var starDate string
	var endDate string

	type tweetData struct {
		Total int     `josn:"total"`
		Num   string  `json:"unique_tags"`
		Ratio float32 `json:"ratio"`
		Party string  `json:"name"`
		Tag   string  `json:"tag"`
	}

	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
	}

	var tweets []tweetData
	err = json.Unmarshal(data, &tweets)
	if err != nil {
		log.Println("Json unmarshal error: \n", err)
	}


	type Days struct {
		Ratio [][]float32 `json:"ratio"`
	}

	type RequestedInterval struct {
		Ratio []float32 `json:"ratio"`
	}

	type TweetParty struct {
		Name              string            `json:"name"`
		Limit             int               `json:"limit"`
		StartDate         string            `json:"startDate"`
		EndDate           string            `json:"endDate"`
		Hashtags          []string          `json:"hastags"`
		UniqueTags        int               `json:"uniqueTags"`
		Days              Days              `json:"days"`
		RequestedInterval RequestedInterval `json:"requestedInterval"`
	}

	var send TweetParty

	send.Name = tweets[0].Party
	send.Limit = limit
	send.StartDate = starDate
	send.EndDate = endDate



	var hastags []string
	var ratioDays [][]float32
	var ratioTotal []float32

	for i := 0;i < limit;i++{
		var days []float32
		days = append(days,tweets[i].Ratio)
		ratioDays = append(ratioDays,days)
		ratioTotal = append(ratioTotal,tweets[i].Ratio)
		hastags = append(hastags,tweets[i].Tag)
	}


	send.Hashtags = hastags
	send.UniqueTags = len(tweets)
	send.Days = Days{ratioDays}
	send.RequestedInterval = RequestedInterval{ratioTotal}

	json.NewEncoder(w).Encode(send)
}
