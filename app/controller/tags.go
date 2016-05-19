package controller

import (
	"encoding/json"
	"lcd/PTC-search-service/app/models"
	"lcd/PTC-search-service/app/web"
	"net/http"

	"errors"
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

	//Check Parameters from the request
	accounts, err := web.Param(r, models.TagsParameter(models.GROUP))
	if err != nil {
		writeError(w, err)
		return
	}
	starDate, err := web.Param(r, models.TagsParameter(models.START_DATE))
	if err != nil {
		writeError(w, err)
		return
	}
	endDate, err := web.Param(r, models.TagsParameter(models.END_DATE))
	if err != nil {
		writeError(w, err)
		return
	}

	limitStr, err := web.Param(r, models.TagsParameter(models.LIMT))
	if err != nil {
		writeError(w, err)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		writeError(w, errors.New("Could not convert limit to string."))
		return
	}

	//Elastic
	accountArray := strings.Split(accounts, ",") // Split into array

	//Set up the response
	var respons models.Tags
	respons.Setup(accounts, starDate, endDate, limit)
	err = respons.CalculateRatio(accountArray)
	if err != nil {
		writeError(w, err)
		return
	}

	json.NewEncoder(w).Encode(respons)
}

func writeError(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(struct {
		Msg string `json:"error_msg"`
	}{err.Error()})
}
