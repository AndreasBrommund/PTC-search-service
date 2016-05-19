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
