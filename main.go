package main

import (
	"github.com/LcKoh0830/backend-test/restapi"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. Return a list of Top Posts ordered by their number of Comments.
	r.GET("/api/v1/top_posts", restapi.TopPosts)

	// 2. Search API Create an endpoint that allows a user to filter the comments based on all the available fields
	r.GET("/api/v1/search", restapi.Search)

	r.Run()
}
