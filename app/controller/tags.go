package controller

import (
	"encoding/json"
	"lcd/PTC-search-service/app/models"
	"lcd/PTC-search-service/app/web"
	"net/http"

	"log"
	"strconv"
	"strings"
)

//Tags, controller function returning the top hashtags based on
//the specified parameters in the Get request.
// group: what twitter accounts to query,
// startDate: from what date
// endDate: until this date
// limit: the limit on the top list, e.g 10 would mean the top 10 hashtags.
func Tags(w http.ResponseWriter, r *http.Request) {

	//Parameters from the request

	accounts, err := web.Param(r, "group")
	if err != nil {
		log.Println("Could not fetch param 'group'")
		log.Println(err)
	}
	starDate, err := web.Param(r, "startDate")
	if err != nil {
		log.Println("Could not fetch param 'startDate'")
		log.Println(err)
	}
	endDate, err := web.Param(r, "endDate")
	if err != nil {
		log.Println("Could not fetch param 'endDate'")
		log.Println(err)
	}

	limitStr, err := web.Param(r, "limit")
	if err != nil {
		log.Println("Could not fetch param 'limit'")
		log.Println(err)
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Println("Could not fetch param 'limit'")
		log.Println(err)
	}

	//Elastic
	accountArray := strings.Split(accounts, ",") // Split into array

	//Set up the response
	var respons models.Tags
	respons.Setup(accounts, starDate, endDate, limit)
	respons.CalculateRatio(accountArray)

	json.NewEncoder(w).Encode(respons)
}
