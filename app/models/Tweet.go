package models

import (
	"gopkg.in/olivere/elastic.v3"
)

type Retweet struct {
	RetweetID string `json:"retweet_id"`
	CreatorID string `json:"creator_id"`
}

type Tweet struct {
	UserID    string   `json:"user_id"`
	TweetID   string   `json:"tweet_id"`
	Text      string   `json:"text"`
	Hashtags  []string `json:"hashtags"`
	Mentions  []string `json:"mentions"`
	Langugage string   `json:"lang"`
	//Retweets  []Retweet             `json:"RT"`
	Date    string                `json:"date"`
	Suggest *elastic.SuggestField `json:"suggest_field,omitempty"`
}
