package controller

import (
	"encoding/json"
	"lcd/PTC-search-service/app/models"
	"lcd/PTC-search-service/app/storage"
	"lcd/PTC-search-service/app/web"
	"net/http"

	"log"
	"strconv"
	"strings"
)

func GetHashtags(w http.ResponseWriter, r *http.Request) {

	//Parameters from the request

	accounts, err := web.Param(r, "group")
	if err != nil {
		log.Println("Could not fetch param 'group'")
		log.Println(err)
	}
	starDate, err := web.Param(r, "starDate")
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
	searchResult := storage.ElasticSearch.GetHashtags(accountArray, starDate, endDate, limit)


	//Set up the response
	var respons models.HashtagData
	respons.Setup(accounts, starDate, endDate, limit)
	respons.CalculateRatio(searchResult)

	json.NewEncoder(w).Encode(respons)
}
