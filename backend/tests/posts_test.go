package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"simple-reddit/common"
	"simple-reddit/posts"
	t_utils "simple-reddit/test_utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreatePost(t *testing.T) {
	var resp = common.APIResponse{}
	commID, err := primitive.ObjectIDFromHex("622034f897a3bc4ddc6333cd")
	if err != nil {
		log.Fatal(err.Error())
	}
	var body = posts.CreatePostRequest{
		UserName:    "johndoe",
		CommunityID: commID,
		Title:       "alpha post!!",
		Body:        "this really an alpha post!!",
	}
	var expected = common.APIResponse{
		Status:  http.StatusCreated,
		Message: common.API_SUCCESS,
		Data:    map[string]interface{}{"data": ""},
	}

	req, err := t_utils.MakeRequest(t_utils.POST, posts.POST_ROUTE_PREFIX, body)
	if err != nil {
		log.Fatal(err.Error())
	}

	// add the customed headers
	for key, value := range customHeaders {
		req.Header.Add(key, value)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(resp)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, expected.Status, resp.Status)
	assert.Equal(t, expected.Message, resp.Message)
}

func TestGetPost(t *testing.T) {
	var resp = common.APIResponse{}
	commID, err := primitive.ObjectIDFromHex("622034f897a3bc4ddc6333cd")
	if err != nil {
		log.Fatal(err.Error())
	}
	var body = posts.CreatePostRequest{
		UserName:    "johndoe",
		CommunityID: commID,
	}
	var expected = common.APIResponse{
		Status:  http.StatusOK,
		Message: common.API_SUCCESS,
		Data:    map[string]interface{}{"data": ""},
	}

	req, err := t_utils.MakeRequest(t_utils.GET, posts.POST_ROUTE_PREFIX, body)
	if err != nil {
		return
	}

	// add the customed headers
	for key, value := range customHeaders {
		req.Header.Add(key, value)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		log.Println(err.Error())
	}

	// log.Println(resp)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected.Status, resp.Status)
	assert.Equal(t, expected.Message, resp.Message)
}
