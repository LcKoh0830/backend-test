package restapi

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/LcKoh0830/backend-test/commentrepo"
	"github.com/LcKoh0830/backend-test/model"
	"github.com/gin-gonic/gin"
)

type fnMatch = func(model.Comment, interface{}) bool

func Search(c *gin.Context) {
	postID, err := GetQueryInt(c.Request.URL.Query(), "_post_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid PostID")
		return
	}
	id, err := GetQueryInt(c.Request.URL.Query(), "_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}
	name := c.Query("_name")
	email := c.Query("_email")
	body := c.Query("_body")

	// Nothing to search
	if postID == nil && id == nil && name == "" && email == "" && body == "" {
		c.JSON(http.StatusOK, []model.Comment{})
		return
	}

	fnMatches := []fnMatch{}
	searchValue := []interface{}{}

	if name != "" {
		fnMatches = append(fnMatches, isNameMatch)
		searchValue = append(searchValue, name)
	}
	if email != "" {
		fnMatches = append(fnMatches, isEmailMatch)
		searchValue = append(searchValue, email)
	}
	if body != "" {
		fnMatches = append(fnMatches, isBodyMatch)
		searchValue = append(searchValue, body)
	}

	comments, err := commentrepo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var result []model.Comment

	// If comment's id is provided, return the comment instead.
	if id != nil {
		c.JSON(http.StatusOK, comments[*id-1])
		return
	}

	// if post's ID is provided, search throught the comment in the post given.
	if postID != nil {
		var temp []model.Comment
		for _, comment := range comments {
			if comment.PostID == *postID {
				temp = append(temp, comment)
			}
		}
		comments = temp
	}

	// if no _name, _email, _body filter provided, return comments with postID given
	// no need to do further filter
	if len(fnMatches) == 0 {
		c.JSON(http.StatusOK, comments)
		return
	}

	for _, comment := range comments {
		for i, isMatch := range fnMatches {
			if isMatch(comment, searchValue[i]) {
				result = append(result, comment)
				break
			}
		}
	}

	if len(result) == 0 {
		result = []model.Comment{}
	}

	c.JSON(http.StatusOK, result)
}

func isNameMatch(comment model.Comment, search interface{}) bool {
	return strings.Contains(comment.Name, search.(string))
}

func isEmailMatch(comment model.Comment, search interface{}) bool {
	return strings.Contains(comment.Email, search.(string))
}

func isBodyMatch(comment model.Comment, search interface{}) bool {
	return strings.Contains(comment.Body, search.(string))
}

func GetQueryInt(query url.Values, paramName string) (*int, error) {
	val := query.Get(paramName)
	if val == "" {
		return nil, nil
	}
	var ival int
	var ival64 int64
	var err error
	ival64, err = strconv.ParseInt(val, 10, 32)
	if err != nil {
		return nil, err
	}
	ival = int(ival64)
	return &ival, nil
}
