package main

import (
	"blog-aggregator/internal/database"
	"context"
	"database/sql"
	"github.com/google/uuid"
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
		type RssFeedModel struct {
			feedXML RssFeedXml
			feedId  uuid.UUID
		}
		feeds := make([]RssFeedModel, 0, 0)
		//fetch
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

				model := RssFeedModel{
					feedXML: feed,
					feedId:  v.ID,
				}
				feeds = append(feeds, model)
			}(v)

		}
		wg.Wait()
		//write to db
		for _, c := range feeds {
			posts := c.feedXML.GetItems()
			for _, post := range posts {
				exist, err := cfg.DB.CheckPostExists(ctx, post.Link)
				if err != nil {
					log.Println(err)
					continue
				}
				if exist {
					continue
				}
				cfg.DB.CreatePost(ctx, database.CreatePostParams{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					Title:     post.Title,
					Url:       post.Link,
					Description: sql.NullString{
						String: post.Description,
						Valid:  post.Description != "",
					},
					PublishedAt: sql.NullTime{
						Time:  parseXMLTime(post.PubDate),
						Valid: true,
					},
					FeedID: c.feedId,
				})
			}
		}

		time.Sleep(time.Second * time.Duration(seconds))
	}

}

func parseXMLTime(s string) time.Time {

	layout := "Tue, 05 Sep 2023 08:00:00 -0000"
	pubDate, _ := time.Parse(layout, s)
	return pubDate
}
