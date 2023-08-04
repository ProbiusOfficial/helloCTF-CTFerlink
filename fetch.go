package main

import (
	"github.com/mmcdole/gofeed"
	"net/http"
)

func (friend *Friend) FetchFeed() error {
	resp, err := http.Get(friend.RSS)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	friend.Feed, err = gofeed.NewParser().Parse(resp.Body)
	if err != nil {
		return err
	}
	//
	//fmt.Println(friend.Feed)
	//fmt.Println(friend.Feed.Title)

	return nil
}
