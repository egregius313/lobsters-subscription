package lobsters

import (
	"io"
	"time"

	"github.com/egregius313/lobsters-subscription/rss"
)

type Post struct {
	Title        string
	ArticleLink  string
	Author       string
	PostDate     time.Time
	CommentsLink string
	Description  string
	Categories   []string
}

func DecodePosts(reader io.Reader) ([]Post, error) {
	items, err := rss.DecodeItems(reader)
	if err != nil {
		return nil, err
	}

	posts := make([]Post, len(items))
	for i, item := range items {
		post := Post{
			Title:        item.Title,
			ArticleLink:  item.Link,
			Author:       item.Author,
			PostDate:     time.Time(item.PubDate),
			CommentsLink: item.Comments,
			Description:  item.Description,
			Categories:   item.Categories,
		}

		posts[i] = post
	}

	return posts, err
}
