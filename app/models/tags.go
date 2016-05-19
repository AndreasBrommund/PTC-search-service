package models

import (
	"lcd/PTC-search-service/app/storage"
)

//Tags a model containg the
//top 'limit' hashtags.
type Tags struct {
	Name      string    `json:"name"`
	Limit     int       `json:"limit"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	Hashtags  []string  `json:"hashtags"`
	Ratio     []float32 `json:"ratio"`
}

//Setup inits a new Tags pointer.
func (this *Tags) Setup(name, start, end string, limit int) {
	this.Name = name
	this.StartDate = start
	this.EndDate = end
	this.Limit = limit
}

//CalculateRatio is responsbile for making requests to elastic and
//performing ratio calculates on the returned data. Then populating the
//Hashtags and Ratio arrays.
func (this *Tags) CalculateRatio(accountArray []string) {
	searchResult := storage.ElasticSearch.GetHashtags(accountArray,
		this.StartDate, this.EndDate, this.Limit)
	topTag, _ := searchResult.Aggregations.Terms("top_tags")
	total := topTag.SumOfOtherDocCount //The total numbers of hashtags except the top (limit) hashtags
	var hashtags []string
	var ratio []float32

	for _, d := range topTag.Buckets {
		hashtags = append(hashtags, d.Key.(string))
		ratio = append(ratio, float32(d.DocCount)) //It is the total num of hashtags not the ratio
		total += d.DocCount                        //Add the rest of the hashtags to to the total sum
	}

	for i, d := range ratio {
		ratio[i] = d / float32(total) //Calculate the ratio
	}

	this.Hashtags = hashtags
	this.Ratio = ratio
}
