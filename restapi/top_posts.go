package restapi

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/LcKoh0830/backend-test/commentrepo"
	"github.com/LcKoh0830/backend-test/postrepo"
	"github.com/gin-gonic/gin"
)

type Response struct {
	PostID                int    `json:"post_id"`
	PostTitle             string `json:"post_title"`
	PostBody              string `json:"post_body"`
	TotalNumberOfComments int    `json:"total_number_of_comments"`
}

func TopPosts(c *gin.Context) {
	comments, err := commentrepo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	postCommentCountList := map[int]int{}
	for _, comment := range comments {
		if _, ok := postCommentCountList[comment.PostID]; !ok {
			postCommentCountList[comment.PostID] = 1
		} else {
			postCommentCountList[comment.PostID]++
		}
	}
	var result []Response
	for postID, totalComment := range postCommentCountList {
		post, err := postrepo.Load(postID)
		if err != nil {
			log.Fatalf("error to get post: %+v", err)
			continue
		}
		result = append(result, Response{
			PostID:                postID,
			PostTitle:             post.Title,
			PostBody:              post.Body,
			TotalNumberOfComments: totalComment,
		})
	}
	sortByComments(result)
	c.JSON(http.StatusOK, result)
}

func sortByComments(response []Response) []Response {
	if len(response) < 2 {
		return response
	}

	left, right := 0, len(response)-1

	pivot := rand.Int() % len(response)

	response[pivot], response[right] = response[right], response[pivot]

	for i, _ := range response {
		if response[i].TotalNumberOfComments < response[right].TotalNumberOfComments {
			response[left], response[i] = response[i], response[left]
			left++
		}
	}

	response[left], response[right] = response[right], response[left]

	sortByComments(response[:left])
	sortByComments(response[left+1:])

	return response
}
