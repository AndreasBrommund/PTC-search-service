package storage

import (
	"log"

	"gopkg.in/olivere/elastic.v3"
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

func (this *Elastic) GetHastags(twitter_id string, from string, to string, limit int) *elastic.SearchResult {
	termQuery := elastic.NewTermQuery("following", twitter_id)
	topTagsAgg := elastic.NewTermsAggregation().Field("hashtags").Size(limit)
	searchResult, err := this.client.Search().
		Index("tweets").                     // search in index "twitter"
		Query(termQuery).                    // specify the query
		From(0).Size(0).                     // take documents 0-9
		Pretty(true).                        // pretty print request and response JSON
		Aggregation("top_tags", topTagsAgg). //Agg func
		Do()                                 // execute
	if err != nil {
		//Fix
		panic(err)
	}
	return searchResult
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
