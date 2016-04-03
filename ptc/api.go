package ptc

import (
	"encoding/json"
	"net/http"

)

//apiVersion is a simple endpoint handler that
//just writes a simple struct as json as the HTTP response.
func apiVersion(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		}{"PTC-Search-Service", "0.1"})
}

// returns first and last tweet from a specific party
// in a json object.
func tweetRange(w http.ResponseWriter, r *http.Request) {
	
	type party struct {
		name 	string
	}

	p := party{"socialdemmokraterna"}
	
	//party will be sent to elastic and hopfully we can 
	//retrive first and last post from that party
	//and then return
	json.NewEncoder(w).Encode(
		struct {
			Party		string `json:"party"`
			MinDate   	string `json:"minDate"`
			MaxDate 	string `json:"maxDate"`
		}{p.name,"2006-02-01T00:00:00.000Z", "2016-02-01T00:00:00.000Z"})
}
