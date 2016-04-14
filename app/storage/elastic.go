package storage

import (
	"log"

	"encoding/json"
	"gopkg.in/olivere/elastic.v3"
	"lcd/PTC-search-service/app/models"
	"net/http"
	"reflect"
)

var ElasticSearch Elastic

// A type that holds our Elastic client, to prevent us from creating it over and over again
type Elastic struct {
	client *elastic.Client
}

// Create the elastic client and return a Elastic struct containing the client
func Connect() error {
	// Create the Elasticsearch client
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		log.Println("Could not create Elasticsearch client")
		log.Println(err)
		return err
	}
	ElasticSearch = Elastic{client}
	return nil
}

// A (temporary) search method that returns an elastic SearchResult object that can be looped over
func (this *Elastic) SearchTweetsFromID(user_id string) *elastic.SearchResult {

	// Make a search
	//
	//
	//Aggregations to be used later:
	//dateRangeAgg := NewDateRangeAggregation().Field("created").Lt("2012-01-01").Between("2012-01-01", "2013-01-01").Gt("2013-01-01")
	//https://github.com/olivere/elastic/blob/release-branch.v3/search_aggs_test.go
	//termQuery := elastic.NewTermQuery("user_id", user_id)
	topTagsHitsAgg := elastic.NewTopHitsAggregation().Sort("user_id", true).Size(5).FetchSource(true)
	topTagsAgg := elastic.NewTermsAggregation().Field("hashtags").Size(3).SubAggregation("top_tag_hits", topTagsHitsAgg)
	searchResult, err := this.client.Search().
		Index("test-index").               // search in index "twitter"
		Query(elastic.NewMatchAllQuery()). // specify the query
		Sort("user_id", true).             // sort by "user" field, ascending
		From(0).Size(10000).               // take documents 0-9
		Pretty(true).                      // pretty print request and response JSON
		Aggregation("top-tags", topTagsAgg).
		Do() // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	return searchResult

}

// A function that returns the latest or earliest date for tweet from given account
func (this *Elastic) GetTweetDateForUser(user_id string, w http.ResponseWriter) {
	// Our return values
	var firstDate, lastDate string

	// Search for the earliest tweet from a person that is following the given user_id
	termQuery := elastic.NewTermQuery("following", user_id)
	searchResult, err := this.client.Search().
		Index("tweets").    // search in index "tweets"
		Query(termQuery).   // specify the query
		Sort("date", true). // sort by "date" field, ascending ==> true
		From(0).Size(1).    // take document 0
		Pretty(true).       // pretty print request and response JSON

		Do() // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	var ttyp models.Tweet
	// Loop through the results
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(models.Tweet); ok {
			firstDate = t.Date
		}
	}
	// Search for the latest tweet from a person that is following the given user_id
	searchResult, err = this.client.Search().
		Index("tweets").     // search in index "twitter"
		Query(termQuery).    // specify the query
		Sort("date", false). // sort by "date" field, false => descending
		From(0).Size(1).     // take document 0
		Pretty(true).        // pretty print request and response JSON

		Do() // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	// Loop through the result
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(models.Tweet); ok {
			lastDate = t.Date
		}
	}

	// Format JSON and write to response
	json.NewEncoder(w).Encode(
		struct {
			StartDate string `json:"startDate"`
			EndDate   string `json:"endDate"`
		}{firstDate, lastDate})
	return

}
