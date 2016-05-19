package models

import (
	"lcd/PTC-search-service/app/storage"
	"reflect"
)

//Interval is a model of the data representing start and endDate over a
//dataset.
type Interval struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

//GetInterval populate the fields StartDate and EndDate using the
//dateSearch function.
func (this *Interval) GetInterval(twitter_id string) {
	this.StartDate = this.dateSearch(twitter_id, true)
	this.EndDate = this.dateSearch(twitter_id, false)
}

//dataSearch asks elastic for the tweet based on the account id and pulls the first tweet
//of the specifed order, asc = true means that the newst will be returned.
//will then parse the tweet and return the date.
func (this *Interval) dateSearch(twitter_id string, asc bool) (date string) {
	// Search for the earliest tweet from a person that is following the given user_id
	searchResult := storage.ElasticSearch.GetDateInterval(twitter_id, asc)
	var tweet Tweet
	// Loop through the results
	for _, item := range searchResult.Each(reflect.TypeOf(tweet)) {
		if t, ok := item.(Tweet); ok {
			date = t.Date
		}
	}
	return
}
