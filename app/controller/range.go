package controller

import (
	"encoding/json"
	"lcd/PTC-search-service/app/models"
	"lcd/PTC-search-service/app/web"
	"log"
	"net/http"
)

// returns first and last tweet from a specific party
// in a json object.
func TweetRange(w http.ResponseWriter, r *http.Request) {

	type account struct {
		id string
	}

	accountId, err := web.Param(r, "account")
	if err != nil {
		log.Println("Could not fetch param 'account'")
		log.Println(err)
	}

	//acc := account{accountID}

	//storage.ElasticSearch.GetTweetDateForUser(accountID, w)

	var rng = models.Range{}
	rng.GetRange(accountId)
	json.NewEncoder(w).Encode(rng)

	//party will be sent to elastic and hopfully we can
	//retrive first and last post from that party
	//and then return
	/*json.NewEncoder(w).Encode(
	struct {
		Party   string `json:"party"`
		MinDate string `json:"minDate"`
		MaxDate string `json:"maxDate"`
	}{p.name, "2006-02-01T00:00:00.000Z", "2016-02-01T00:00:00.000Z"})*/
}
