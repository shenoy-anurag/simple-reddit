package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"simple-reddit/common"
	t_utils "simple-reddit/test_utils"
	"simple-reddit/posts"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	var resp = common.APIResponse{}
  commID, err := primitive.ObjectIDFromHex("622034f897a3bc4ddc6333cd")
	var body = posts.CreatePostRequest{
    UserName: "johndoe",
	  CommunityID: commID,
	   Title : "alpha post!!",
    Body : "this really an alpha post!!",
}
	// isoDateTime, err := time.Parse("YYYY-mm-ddTHH:MM:ssZ", "2022-01-21T04:41:56.616Z")
	// var user = users.CreateUserRequest{
	// 	FirstName: "Albert",
	// 	LastName:  "Einstein",
	// 	Email:     "einsteinalbert@ufl.edu",
	// 	Username:  "albert",
	// 	Password:  "$2a$10$QYXnNZJH.hNuNOiT8Tq8nOJy02V0mcyT5h9ARQvA2bO35rdd72Zym",
	// }
	// users.CreateUserInDB(user)
	// userDB, _ := users.GetUserDetails(body.Username)
	// userResp := users.ConvertUserDBModelToUserResponse(userDB)
	var expected = common.APIResponse{
    Status:  http.StatusCreated,
    Message: common.API_SUCCESS,
    Data:    map[string]interface{}{"data": ""},
}

	req, err := t_utils.MakeRequest(t_utils.POST, posts.POST_ROUTE_PREFIX, body)
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

	log.Println(resp)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, expected.Status, resp.Status)
	assert.Equal(t, expected.Message, resp.Message)
}

func TestGetPost(t *testing.T) {
	var resp = common.APIResponse{}
  commID, err := primitive.ObjectIDFromHex("622034f897a3bc4ddc6333cd")
	var body = posts.CreatePostRequest{
    UserName: "johndoe",
	  CommunityID: commID,
}
	// isoDateTime, err := time.Parse("YYYY-mm-ddTHH:MM:ssZ", "2022-01-21T04:41:56.616Z")
	// var user = users.CreateUserRequest{
	// 	FirstName: "Albert",
	// 	LastName:  "Einstein",
	// 	Email:     "einsteinalbert@ufl.edu",
	// 	Username:  "albert",
	// 	Password:  "$2a$10$QYXnNZJH.hNuNOiT8Tq8nOJy02V0mcyT5h9ARQvA2bO35rdd72Zym",
	// }
	// users.CreateUserInDB(user)
	// userDB, _ := users.GetUserDetails(body.Username)
	// userResp := users.ConvertUserDBModelToUserResponse(userDB)
	var expected = common.APIResponse{
    Status:  http.StatusCreated,
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

	log.Println(resp)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected.Status, resp.Status)
	assert.Equal(t, expected.Message, resp.Message)
}
