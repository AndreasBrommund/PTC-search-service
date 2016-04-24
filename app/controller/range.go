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
	accountId, err := web.Param(r, "account")
	if err != nil {
		log.Println("No 'account' specified. Date range will be for whole dateset.")
	}

	var rng = models.Range{}
	rng.GetRange(accountId)
	json.NewEncoder(w).Encode(rng)
}
