package models

type HashtagData struct {
	Name      string    `json:"name"`
	Limit     int       `json:"limit"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	Hashtags  []string  `json:"hastags"`
	Ratio     []float32 `json:"ratio"`
}
