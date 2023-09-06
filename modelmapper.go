package main

import (
	"blog-aggregator/internal/database"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Apikey    string    `json:"apikey"`
}

func dbUserToUser(dbUser database.User) (user User) {
	user.ID = dbUser.ID
	user.CreatedAt = dbUser.CreatedAt
	user.UpdatedAt = dbUser.UpdatedAt
	user.Name = dbUser.Name
	user.Apikey = dbUser.Apikey
	return
}

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	UserID        uuid.UUID  `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func dbFeedToFeed(dbFeed database.Feed) (feed Feed) {
	feed.ID = dbFeed.ID
	feed.CreatedAt = dbFeed.CreatedAt
	feed.UpdatedAt = dbFeed.UpdatedAt
	feed.Name = dbFeed.Name
	feed.Url = dbFeed.Url
	feed.UserID = dbFeed.UserID
	if dbFeed.LastFetchedAt.Valid {
		*feed.LastFetchedAt = dbFeed.LastFetchedAt.Time
	} else {
		feed.LastFetchedAt = nil
	}
	return
}
