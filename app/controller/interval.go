package controller

import (
	"encoding/json"
	"lcd/PTC-search-service/app/models"
	"lcd/PTC-search-service/app/web"
	"log"
	"net/http"
)

//Interval, controller function of the /interval route.
// Returns first and last tweet from a specific party or the entire dataset.
// This is decided based on if the 'account'
// parameters is passed to the request or not.
func Interval(w http.ResponseWriter, r *http.Request) {
	accountId, err := web.Param(r, "account")
	if err != nil {
		log.Println("No 'account' specified. Date range will be for whole dateset.")
	}

	var interval = models.Interval{}
	interval.GetInterval(accountId)
	json.NewEncoder(w).Encode(interval)
}
