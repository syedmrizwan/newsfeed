package handler

import (
	"fmt"
	"net/http"
	"newsfeeder/platform/newsfeed"
	"newsfeeder/util"

	"github.com/gin-gonic/gin"
	// "github.com/tidwall/gjson"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
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
