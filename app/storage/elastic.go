package storage

import (
	"log"

	"gopkg.in/olivere/elastic.v3"
)

var ElasticSearch Elastic

// A type that holds our Elastic client, to prevent us from creating it over and over again
type Elastic struct {
	Client *elastic.Client
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

func (this *Elastic) GetHashtags(twitter_id string, from string, to string, limit int) *elastic.SearchResult {
	termQuery := elastic.NewTermQuery("following", twitter_id)    //Get tweets from  the right account
	rangeQuery := elastic.NewRangeQuery("date").From(from).To(to) //Get tweets in the time interval
	boolQuery := elastic.NewBoolQuery().Must(termQuery, rangeQuery)
	topTagsAgg := elastic.NewTermsAggregation().Field("hashtags").Size(limit)
	searchResult, err := this.Client.Search().
		Index("tweets").  // search in index "twitts"
		Query(boolQuery). // specify the query
		From(0).Size(0).
		Pretty(true).                        // pretty print request and response JSON
		Aggregation("top_tags", topTagsAgg). //Agg func
		Do()                                 // execute
	if err != nil {
		panic(err)
	}
	return searchResult
}