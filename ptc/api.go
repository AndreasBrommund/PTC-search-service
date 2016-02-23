package ptc

import (
	"encoding/json"
	"net/http"
)

func apiVersion(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		}{"PTC-Search-Service", "0.1"})
}
func countTweet(w http.ResponseWriter, r *http.Request) {
	tweets, err := database.GetNumberOfTweets()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(tweets)
}
