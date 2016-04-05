package models

import (
	"time"

	"gopkg.in/olivere/elastic.v3"
)

type Tweet struct {
	User     string                `json:"user_id"`
	Message  string                `json:"text"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Hashtags []string              `json:"hashtags"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}
