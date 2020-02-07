package handler

import (
	"net/http"
	"newsfeeder/database"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetConnection()

		requestBody := newsfeedPostRequest{}
		c.Bind(&requestBody)

		item := &newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		err := db.Insert(item)
		if err != nil {
			panic(err)
		}
		// feed.Add(item)
		c.Status(http.StatusNoContent)
	}

}
