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
	
	
	
	json.NewEncoder(w).Encode(
		struct {
			StartDate   string `json:""`
			EndDate 	string `json:""`
		}{"PTC-Search-Service", "0.1"})
}
