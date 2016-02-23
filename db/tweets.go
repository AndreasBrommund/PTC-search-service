package db

import "github.com/lcd/PTC-search-service/model"

func (this *Database) GetNumberOfTweets() (tweet model.Tweets, err error) {
	rows, err := this.conn.Query("SELECT COUNT(*) as count FROM tweet")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&tweet.Nr)
	}
	return tweet, nil
}
