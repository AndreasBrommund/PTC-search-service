package storage

import "gopkg.in/olivere/elastic.v3"

//GetDateInterval builds a query asking for
//the specified twitter id's news or oldest tweet based on the
//asc parameter. asc true means that news will be on top.
func (this *Elastic) GetDateInterval(twitter_id string, asc bool) *elastic.SearchResult {

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
