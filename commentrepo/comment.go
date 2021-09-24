package commentrepo

import (
	"encoding/json"
	"log"

	"github.com/LcKoh0830/backend-test/model"
	"github.com/go-resty/resty/v2"
)

const COMMENT_ENDPOINT = "https://jsonplaceholder.typicode.com/comments"

func List() ([]model.Comment, error) {
	client := resty.New()

	resp, err := client.R().Get(COMMENT_ENDPOINT)
	if err != nil {
		return nil, err
	}
	var comments []model.Comment
	if err = json.Unmarshal(resp.Body(), &comments); err != nil {
		log.Fatalf("Failed to unmarshal: %+v", err)
		return nil, err
	}
	return comments, nil
}
