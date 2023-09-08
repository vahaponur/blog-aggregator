package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFeed(t *testing.T) {
	feed, err := GetFeed(url.URL{RawPath: "https://rss.art19.com/apology-line"})

	require.NoError(t, err)
	require.NotEmpty(t, feed)
}
