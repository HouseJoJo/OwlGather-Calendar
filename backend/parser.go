package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func test() {
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL("https://owllife.kennesaw.edu/events.rss")
	if err != nil {
		fmt.Println("Error parsing RSS feed:", err)
		return
	}

	if len(feed.Items) == 0 {
		fmt.Println("No items found in the RSS feed.")
		return
	}

	for _, item := range feed.Items {
		fmt.Println("Title: ", item.Title)
	}

}
