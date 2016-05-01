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

func (this *Elastic) GetHashtags(twitter_ids []string, from string, to string, limit int) *elastic.SearchResult {
	// Loop through the accounts and create termsBoolQuery, a query with the 'should' argument so that we must have any of the given accounts for value following
	termsBoolQuery := elastic.NewBoolQuery()
	for _, account := range twitter_ids {
		termQuery := elastic.NewTermQuery("following", account) //Get tweets from  the right account
		termsBoolQuery = termsBoolQuery.Should(termQuery)
	}
	rangeQuery := elastic.NewRangeQuery("date").From(from).To(to)        //Get tweets in the time interval
	boolQuery := elastic.NewBoolQuery().Must(termsBoolQuery, rangeQuery) // Must match some of the 'should' clauses and the range
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
