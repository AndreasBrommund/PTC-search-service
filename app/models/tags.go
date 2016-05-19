package models

import (
	"lcd/PTC-search-service/app/storage"
	"reflect"
)

const (
	GROUP = iota
	LIMT
	START_DATE
	END_DATE
	HASHTAGS
	RATIO
)

//Tags a model containg the
//top 'limit' hashtags.
type Tags struct {
	Group     string    `json:"group"`
	Limit     int       `json:"limit"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	Hashtags  []string  `json:"hashtags"`
	Ratio     []float32 `json:"ratio"`
}

func TagsParameter(index int) string {
	return reflect.TypeOf(Tags{}).Field(index).Tag.Get("json")
}

//Setup inits a new Tags pointer.
func (this *Tags) Setup(group, start, end string, limit int) {
	this.Group = group
	this.StartDate = start
	this.EndDate = end
	this.Limit = limit
}

//CalculateRatio is responsbile for making requests to elastic and
//performing ratio calculates on the returned data. Then populating the
//Hashtags and Ratio arrays.
func (this *Tags) CalculateRatio(accountArray []string) error {
	searchResult, err := storage.ElasticSearch.GetHashtags(accountArray,
		this.StartDate, this.EndDate, this.Limit)

	if err != nil {
		return err
	}

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
	return err
}
