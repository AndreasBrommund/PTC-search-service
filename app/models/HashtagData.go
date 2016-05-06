package models

import "gopkg.in/olivere/elastic.v3"

type HashtagData struct {
	Name      string    `json:"name"`
	Limit     int       `json:"limit"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	Hashtags  []string  `json:"hashtags"`
	Ratio     []float32 `json:"ratio"`
}

func (this *HashtagData) Setup(name, start, end string, limit int) {
	this.Name = name
	this.StartDate = start
	this.EndDate = end
	this.Limit = limit
}

func (this *HashtagData) CalculateRatio(searchResult *elastic.SearchResult) {
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
