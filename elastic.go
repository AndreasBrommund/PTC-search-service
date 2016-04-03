package main

import (
    "fmt"
    "log"
    "gopkg.in/olivere/elastic.v3"
    "time"
    "reflect"
)

type Tweet struct {
  User     string                `json:"user_id"`
  Message  string                `json:"text"`
  Retweets int                   `json:"retweets"`
  Image    string                `json:"image,omitempty"`
  Created  time.Time             `json:"created,omitempty"`
  Tags     []string              `json:"tags,omitempty"`
  Location string                `json:"location,omitempty"`
  Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

func main() {
    // Create the Elasticsearch client
    client, err := elastic.NewClient()
    if err != nil {
        // Handle error
        log.Println("Could not create Elasticsearch client")
        log.Println(err)
        panic(err)
    }

    // Make a search
    termQuery := elastic.NewTermQuery("user_id", "100004471")
    searchResult, err := client.Search().
        Index("test-index").   // search in index "twitter"
        Query(termQuery).   // specify the query
        Sort("user_id", true). // sort by "user" field, ascending
        From(0).Size(10).   // take documents 0-9
        Pretty(true).       // pretty print request and response JSON
        Do()                // execute
    if err != nil {
        // Handle error
        panic(err)
    }

    // searchResult is of type SearchResult and returns hits, suggestions,
    // and all kinds of other information from Elasticsearch.
    fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

    // Loop through the search results
    var ttyp Tweet
    for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
        if t, ok := item.(Tweet); ok {
            fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
        }
    }
    // Return total hits
    fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

}