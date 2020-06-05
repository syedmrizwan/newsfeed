package handler

import (
	"context"
	"fmt"
	"net/http"
	"newsfeeder/database"
	"newsfeeder/platform/newsfeed"
	"newsfeeder/util"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/tidwall/gjson"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetConnection()
		var items []newsfeed.Item
		if err := db.Model(&items).Select(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func GetBitcoinAndSportsNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := util.AsyncHTTP([]string{"bitcoin", "sports"})
		if err != nil {
			c.JSON(400, "Bad request")
		}
		fmt.Println(resp)
		c.JSON(http.StatusOK, resp)
	}
}

// GetSportsNews godoc
// @Summary Retrieves Sports News
// @Produce json
// @Success 200
// @Router /sportsnews [get]
func GetSportsNews() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create a new context
		// With a deadline of 100 milliseconds
		ctx := context.Background()
		ctx, _ = context.WithTimeout(ctx, 10*time.Second)

		resp, err := util.GetURLResponse(ctx, "sports")
		if err != nil {
			c.JSON(400, "Bad request")
		}
		fmt.Println(resp)
		c.Data(http.StatusOK, "application/json", resp)
	}
}
