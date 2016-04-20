package models

import (
	"lcd/PTC-search-service/app/storage"
	"reflect"
)

type Range struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func (this *Range) GetRange(twitter_id string) {
	this.StartDate = this.dateSearch(twitter_id,true)
	this.EndDate = this.dateSearch(twitter_id,false)
}

func (this *Range) dateSearch(twitter_id string, asc bool) (date string) {
	// Search for the earliest tweet from a person that is following the given user_id
	searchResult := storage.ElasticSearch.GetTweetDate(twitter_id,asc)
	var tweet Tweet
	// Loop through the results
	for _, item := range searchResult.Each(reflect.TypeOf(tweet)) {
		if t, ok := item.(Tweet); ok {
			date = t.Date
		}
	}
	return
}
