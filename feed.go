package main

import (
	"github.com/gorilla/feeds"
)

func createFeed(posts []Post) *feeds.Feed {
	feed := &feeds.Feed{
		Title:       "confusedalex's blog",
		Link:        &feeds.Link{Href: "https://confusealex.dev/blog"},
		Description: "my personal blog",
		Author:      &feeds.Author{Name: "confusedalex", Email: "hello@confusedalex.dev"},
	}

	var items []*feeds.Item

	for _, post := range posts {
		if post.Draft != true {

			// link := templ.SafeURL(path.Join(post.Slug, "/"))

			items = append(items, &feeds.Item{
				Title:   post.Title,
				Created: post.Date,
				// Link:    link,
			})
		}
	}

	feed.Items = items

	return feed
}
