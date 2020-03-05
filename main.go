package main

import (
	"log"
	"net/http"
	"newsfeeder/httpd/handler"
	"newsfeeder/middleware"
	"newsfeeder/platform/newsfeed"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// the jwt middleware
	authMiddleware, err := middleware.SetUpAuth()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/api/v1")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		feed := newsfeed.New()
		auth.GET("/newsfeed", handler.NewsfeedGet(feed))
		auth.GET("/bitcoinnews", handler.GetBitcoinAndSportsNews())
		auth.GET("/sportsnews", handler.GetSportsNews())
		auth.POST("/newsfeed", handler.NewsfeedPost(feed))
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}

}
