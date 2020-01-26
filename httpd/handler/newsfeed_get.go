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

func GetBitcoinNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := util.AsyncHTTP([]string{"bitcoin", "sports"})
		if err != nil {
			c.JSON(400, "Bad request")
		}
		fmt.Println(resp)
		c.JSON(http.StatusOK, resp)
		// url := "https://newsapi.org/v2/everything"
		// params := map[string]string{
		// 	"q":      "bitcoin",
		// 	"from":   "2020-1-15",
		// 	"sortBy": "publishedAt",
		// 	"apiKey": "e4d1a5d882eb439ea2471a6d9948ac1c"}
		// client := resty.New()
		// client.SetQueryParams(params)
		// resp, err := client.R().Get(url)
		// if err != nil {
		// 	fmt.Println(err)
		// 	c.JSON(400, "Bad request")
		// 	return
		// }
		//c.Data(http.StatusOK, "application/json", resp.Body())
	}
}
