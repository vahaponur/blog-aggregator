package main

import (
	"blog-aggregator/internal/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"
)

// n is how many feeds will be getted in one fetch second is the wait time to next fetch
func fetchNext(n int32, seconds int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for {
		ctx := context.Background()
		items, err := cfg.DB.GetNextFeedsToFetch(ctx, n)
		if err != nil {
			log.Fatal(err)
		}
		feeds := make([]RssFeedXml, 0, 0)

		for _, v := range items {
			wg.Add(1)
			rp := v.Url
			go func(v database.Feed) {
				mu.Lock()
				defer mu.Unlock()
				defer wg.Done()

				feed, err := GetFeed(url.URL{RawPath: rp})
				if err != nil {
					log.Println(err)
					return
				}
				cfg.DB.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
					ID: v.ID,
					LastFetchedAt: sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					},
				})
				fmt.Println(feed.Channel.Title)
				feeds = append(feeds, feed)
			}(v)

		}
		wg.Wait()
		time.Sleep(time.Second * time.Duration(seconds))
	}

}
