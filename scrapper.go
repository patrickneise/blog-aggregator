package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/patrickneise/blog-aggregator/internal/database"
)

func startScraping(db *database.Queries, concurrency int, frequency time.Duration) error {
	log.Printf("Collecting feeds every %s on %v goroutines...", frequency, concurrency)

	ticker := time.NewTicker(frequency)

	for ; ; <-ticker.C {

		feeds, err := db.GetNextFeedToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Couldn't retreive feeds to fetch from the DB", err)
			continue
		}
		log.Printf("Fetching %v feeds", len(feeds))
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, feed, wg)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.ID, err)
		return
	}

	feedData, err := fetchFeed(feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	err = processFeedData(feedData)
	if err != nil {
		return
	}
}

func fetchFeed(url string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	decoder := xml.NewDecoder(res.Body)
	feedData := RSSFeed{}
	err = decoder.Decode(&feedData)
	if err != nil {
		return nil, err
	}

	return &feedData, nil
}

func processFeedData(feedData *RSSFeed) error {
	for _, item := range feedData.Channel.Item {
		fmt.Println(item.Title)
	}
	return nil
}
