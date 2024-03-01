package domain

import "net/url"

type RSSFeedLink struct {
	Link        url.URL `json:"link"`
	Version     int8    `json:"version"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
}
