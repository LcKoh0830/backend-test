package postrepo

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/LcKoh0830/backend-test/model"
	"github.com/go-resty/resty/v2"
)

const SINGLE_POST_ENDPOINT = "https://jsonplaceholder.typicode.com/posts/"

func Load(id int) (model.Post, error) {
	client := resty.New()
	var post model.Post
	resp, err := client.R().Get(fmt.Sprintf("%s/%d", SINGLE_POST_ENDPOINT, id))
	if err != nil {
		return post, err
	}
	if err := json.Unmarshal(resp.Body(), &post); err != nil {
		log.Fatalf("failed to unmarshal: %+v", err)
		return post, err
	}
	return post, nil
}
