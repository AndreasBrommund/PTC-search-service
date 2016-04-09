package models

type Days struct {
	Ratio [][]float32 `json:"ratio"`
}

type RequestedInterval struct {
	Ratio []float32 `json:"ratio"`
}

type TweetParty struct {
	Name              string            `json:"name"`
	Limit             int               `json:"limit"`
	StartDate         string            `json:"startDate"`
	EndDate           string            `json:"endDate"`
	Hashtags          []string          `json:"hastags"`
	UniqueTags        int               `json:"uniqueTags"`
	Days              Days              `json:"days"`
	RequestedInterval RequestedInterval `json:"requestedInterval"`
}
