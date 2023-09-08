package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
)

type RssFeedXml struct {
	XMLName          xml.Name `xml:"rss"`
	Version          string   `xml:"version,attr"`
	ContentNamespace string   `xml:"xmlns:content,attr"`
	Channel          *RssChannel
}
type RssChannel struct {
	XMLName     xml.Name   `xml:"channel"`
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	Language    string     `xml:"language,omitempty"`
	Copyright   string     `xml:"copyright,omitempty"`
	Items       []*RssItem `xml:"item"`
}
type RssItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Author      string   `xml:"author,omitempty"`
	PubDate     string   `xml:"pubDate,omitempty"`
}

func GetFeed(link url.URL) (RssFeedXml, error) {
	resp, err := http.Get(link.RawPath)
	if err != nil {
		return RssFeedXml{}, err
	}
	feed := RssFeedXml{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RssFeedXml{}, err
	}
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return RssFeedXml{}, err
	}
	return feed, nil
}
func (feed *RssFeedXml) GetItems() []*RssItem {
	return feed.Channel.Items
}
