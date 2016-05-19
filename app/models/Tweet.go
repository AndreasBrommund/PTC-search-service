package models

import (
	"gopkg.in/olivere/elastic.v3"
)

//Tweet a model struct that represents a typical
//tweet in elastic.
type Tweet struct {
	UserID    string                `json:"user_id"`
	TweetID   string                `json:"tweet_id"`
	Text      string                `json:"text"`
	Hashtags  []string              `json:"hashtags"`
	Mentions  []string              `json:"mentions"`
	Langugage string                `json:"lang"`
	Date      string                `json:"date"`
	Suggest   *elastic.SuggestField `json:"suggest_field,omitempty"`
}
