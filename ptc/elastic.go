package ptc

import (
    "log"
    "gopkg.in/olivere/elastic.v3"
    "time"
)
/*
type Tweet struct {
  User     string                `json:"user_id"`
  Message  string                `json:"text"`
  Retweets int                   `json:"retweets"`
  TweetID    string                `json:"tweet_id"`
  Created  time.Time             `json:"created,omitempty"`
  Hashtags     string              `json:"hashtags,omitempty"`
  Location string                `json:"location,omitempty"`
  Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}*/

type Tweet struct {
    User     string                `json:"user_id"`
    Message  string                `json:"text"`
    Retweets int                   `json:"retweets"`
    Image    string                `json:"image,omitempty"`
    Created  time.Time             `json:"created,omitempty"`
    Tags     []string              `json:"tags,omitempty"`
    Location string                `json:"location,omitempty"`
    Hashtags     []string              `json:"hashtags"`
    Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

// A type that holds our Elastic client, to prevent us from creating it over and over again
type Elastic struct {
    client *elastic.Client 
}

// Create the elastic client and return a Elastic struct containing the client
func NewElastic() (Elastic, error) {
    // Create the Elasticsearch client
    client, err := elastic.NewClient()
    if err != nil {
        // Handle error
        log.Println("Could not create Elasticsearch client")
        log.Println(err)
        panic(err)
    }
    return Elastic{client}, err
}

// A (temporary) search method that returns an elastic SearchResult object that can be looped over
func (this *Elastic) SearchTweetsFromID(user_id string) *elastic.SearchResult {
    
    // Make a search
    //termQuery := elastic.NewTermQuery("user_id", user_id)
    searchResult, err := this.client.Search().
        Index("test-index").   // search in index "twitter"
        Query(elastic.NewMatchAllQuery()).   // specify the query
        Sort("user_id", true). // sort by "user" field, ascending
        From(0).Size(10000).   // take documents 0-9
        Pretty(true).       // pretty print request and response JSON
        Do()                // execute
    if err != nil {
        // Handle error
        panic(err)
    }
    return searchResult

}