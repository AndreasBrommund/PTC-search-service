package storage

import "gopkg.in/olivere/elastic.v3"

//GetHashtags queries elastic for the top 'limit' hashtags based from the
//twitter accounts in the twitter_ids array.
func (this *Elastic) GetHashtags(twitter_ids []string, from string, to string, limit int) (*elastic.SearchResult, error) {
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
	return searchResult, err
}
