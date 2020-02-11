package handler

import (
	"net/http"
	"newsfeeder/database"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

type newsfeedPostRequest struct {
	Title string       `json:"title"`
	Post  string       `json:"post"`
	Stats statsRequest `json:"stats"`
}

type statsRequest struct {
	Views int `json:"views"`
	Likes int `json:"likes"`
}

// NewsfeedPost does persist News feed items to database
func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetConnection()

		requestBody := newsfeedPostRequest{}
		c.Bind(&requestBody)

		item := &newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
			Stats: newsfeed.StatsType{
				Views: requestBody.Stats.Views,
				Likes: requestBody.Stats.Likes,
			},
		}
		err := db.Insert(item)
		if err != nil {
			panic(err)
		}
		// feed.Add(item)
		c.Status(http.StatusNoContent)
	}

}
