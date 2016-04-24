package controller

import (
	"encoding/json"
	"lcd/PTC-search-service/app/models"
	"lcd/PTC-search-service/app/storage"
	"lcd/PTC-search-service/app/web"
	"net/http"
	"reflect"

	"gopkg.in/olivere/elastic.v3"

	"log"
	"strconv"
)

func GetHastags(w http.ResponseWriter, r *http.Request) {

	//Parameters from the request

	account, err := web.Param(r, "group")
	if err != nil {
		log.Println("Could not fetch param 'group'")
		log.Println(err)
	}
	starDate, err := web.Param(r, "starDate")
	if err != nil {
		log.Println("Could not fetch param 'startDate'")
		log.Println(err)
	}
	endDate, err := web.Param(r, "endDate")
	if err != nil {
		log.Println("Could not fetch param 'endDate'")
		log.Println(err)
	}

	limitStr, err := web.Param(r, "limit")
	if err != nil {
		log.Println("Could not fetch param 'limit'")
		log.Println(err)
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Println("Could not fetch param 'limit'")
		log.Println(err)
	}

	log.Println(account, starDate, endDate, limit)

	//Elastic
	searchResult := storage.ElasticSearch.GetHastags(account, starDate, endDate, limit)

	res, _ := searchResult.Aggregations.Terms("top_tags")

	//Set up the response
	var respons models.TweetParty

	respons.Name = account
	respons.Limit = limit
	respons.StartDate = starDate
	respons.EndDate = endDate
	total := res.SumOfOtherDocCount //The total numbers of hashtags except the top (limit) hashtags

	var hashtags []string
	var ratio []float32

	for _, d := range res.Buckets {
		hashtags = append(hashtags, d.Key.(string))
		ratio = append(ratio, float32(d.DocCount)) //It is the total num of hashtags not the ratio
		total += d.DocCount                        //Add the rest of the hashtags to to the total sum
	}

	for i, d := range ratio {
		ratio[i] = d / float32(total) //Calculate the ratio
	}

	respons.Hashtags = hashtags
	respons.Ratio = ratio

	json.NewEncoder(w).Encode(respons)
}

func GetTweetsFromUserID(w http.ResponseWriter, r *http.Request) {
	var searchResult *elastic.SearchResult
	// elasticSearch is a global variable defined in server.go containing a Elastic object with a client
	searchResult = storage.ElasticSearch.SearchTweetsFromID("3801501")
	var ttyp models.Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(models.Tweet); ok {
			if len(t.Hashtags) != 0 {
				json.NewEncoder(w).Encode(

					struct {
						User     string   `json:"user_id"`
						Hashtags []string `json:"hashtags"`
					}{t.User, t.Hashtags})
			}
		}
	}
}
