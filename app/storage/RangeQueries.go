package storage

import "gopkg.in/olivere/elastic.v3"

func (this *Elastic) GetDateRange(twitter_id string, asc bool) *elastic.SearchResult {

	if twitter_id == "" {
		searchResult, err := this.Client.Search().
			Index("tweets").   // search in index "tweets"
			Sort("date", asc). // sort by "date" field, ascending ==> true
			From(0).Size(1).   // take document 0
			Pretty(true).      // pretty print request and response JSON
			Do()

		if err != nil {
			panic(err)
		}
		return searchResult
	}

	termQuery := elastic.NewTermQuery("following", twitter_id)
	searchResult, err := this.Client.Search().
		Index("tweets").   // search in index "tweets"
		Query(termQuery).  // specify the query
		Sort("date", asc). // sort by "date" field, ascending ==> true
		From(0).Size(1).   // take document 0
		Pretty(true).      // pretty print request and response JSON
		Do()               // execute

	if err != nil {
		panic(err)
	}

	return searchResult
}
