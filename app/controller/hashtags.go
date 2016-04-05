package controller

import (
	"encoding/json"
	"lcd/PTC-search-service/app/models"
	"reflect"
	"lcd/PTC-search-service/app/storage"
	"gopkg.in/olivere/elastic.v3"
	"net/http"
	"io/ioutil"
	"log"
)



func GetHastags(w http.ResponseWriter, r *http.Request) {

	//Parameters from the request
	var limit int = 100
	var starDate string
	var endDate string


	//This should be replaced by elastic
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

	//Set up the response
	var respons models.TweetParty

	respons.Name = tweets[0].Party
	respons.Limit = limit
	respons.StartDate = starDate
	respons.EndDate = endDate

	var hastags []string
	var ratioDays [][]float32
	var ratioTotal []float32

	for i := 0; i < limit; i++ {
		var days []float32
		days = append(days, tweets[i].Ratio)
		ratioDays = append(ratioDays, days)
		ratioTotal = append(ratioTotal, tweets[i].Ratio)
		hastags = append(hastags, tweets[i].Tag)
	}

	respons.Hashtags = hastags
	respons.UniqueTags = len(tweets)
	respons.Days = models.Days{ratioDays}
	respons.RequestedInterval = models.RequestedInterval{ratioTotal}

	json.NewEncoder(w).Encode(respons)
}

func GetTweetsFromUserID(w http.ResponseWriter, r *http.Request) {
	var searchResult *elastic.SearchResult
	// elasticSearch is a global variable defined in server.go containing a Elastic object with a client
	searchResult = storage.ElasticSearch.SearchTweetsFromID("100004471")
	var ttyp models.Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
        if t, ok := item.(models.Tweet); ok {
        	if len(t.Hashtags) != 0 {
	            json.NewEncoder(w).Encode(
		            struct {
					User    string `json:"user_id"`
					Hashtags 	[]string `json:"hashtags"`
				}{t.User, t.Hashtags})
        	}
        }
    }
}