package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"simple-reddit/comments"
	"simple-reddit/common"
	"simple-reddit/posts"
	t_utils "simple-reddit/test_utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var postID string
var commentID string

const commentBody1 string = "test comment 1"

func CreatePostForComment() string {
	var resp = common.APIResponse{}
	// community we know exists. id corresponds to "maths" community.
	commID, _ := primitive.ObjectIDFromHex("622034f897a3bc4ddc6333cd")
	var body = posts.CreatePostRequest{
		UserName:    "albert",
		CommunityID: commID,
		Title:       "test post title for comment",
		Body:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
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

	postID = resp.Data["data"].(map[string]interface{})["InsertedID"].(string)
	return postID
}

func TestCreateComment(t *testing.T) {
	var resp = common.APIResponse{}
	var expected = common.APIResponse{
		Status:  http.StatusCreated,
		Message: common.API_SUCCESS,
		Data:    map[string]interface{}{"created": map[string]string{"InsertedID": ""}},
	}

	postID := CreatePostForComment()

	body := map[string]interface{}{
		"username":  "albert",
		"post_id":   postID,
		"parent_id": nil,
		"body":      commentBody1,
	}

	req, err := t_utils.MakeRequest(t_utils.POST, comments.COMMENTS_ROUTE_PREFIX, body)
	if err != nil {
		return
	}

	// add the custom headers
	for key, value := range customHeaders {
		req.Header.Add(key, value)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		log.Println(err.Error())
	}

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, expected.Status, resp.Status)
	assert.Equal(t, expected.Message, resp.Message)
	assert.Contains(t, resp.Data, "created")
}

func TestGetComment(t *testing.T) {
	var resp = common.APIResponse{}
	var expected = common.APIResponse{
		Status:  http.StatusOK,
		Message: common.API_SUCCESS,
		Data:    map[string]interface{}{"comments": map[string]interface{}{}},
	}

	queryStr := "?" + "post_id=" + postID

	req, err := t_utils.MakeRequest(t_utils.GET, comments.COMMENTS_ROUTE_PREFIX+queryStr, nil)
	if err != nil {
		return
	}

	// add the custom headers
	for key, value := range customHeaders {
		req.Header.Add(key, value)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		log.Println(err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected.Status, resp.Status)
	assert.Equal(t, expected.Message, resp.Message)
	assert.Contains(t, resp.Data, "comments")

	val := resp.Data["comments"]
	comment := val.([]interface{})[0].(map[string]interface{})
	log.Println("Comment Data:", comment)
	commentMessage := comment["body"]
	assert.Equal(t, commentBody1, commentMessage)
	assert.Equal(t, float64(0), comment["upvotes"])
	commentID = comment["_id"].(string)
}

func TestVoteComment(t *testing.T) {
	var resp = common.APIResponse{}
	var expected = common.APIResponse{
		Status:  http.StatusOK,
		Message: common.API_SUCCESS,
		Data:    map[string]interface{}{"result": "voted"},
	}

	// Upvoting comment created in TestCreateComment()
	body := comments.CommentVoteRequest{
		UserName:  "albert",
		CommentId: commentID,
		Vote:      comments.UPVOTE,
	}

	req, err := t_utils.MakeRequest(t_utils.POST, comments.COMMENTS_ROUTE_VOTE, body)
	if err != nil {
		return
	}

	// add the custom headers
	for key, value := range customHeaders {
		req.Header.Add(key, value)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Vote Response: ", resp.Data)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected.Status, resp.Status)
	assert.Equal(t, expected.Message, resp.Message)
	assert.Contains(t, resp.Data, "result")

	result := resp.Data["result"].(string)
	assert.Equal(t, "voted", result)
	log.Println("All assertions passed for Vote Comment.")
}

func TestDeleteComment(t *testing.T) {
	var resp = common.APIResponse{}
	var expected = common.APIResponse{
		Status:  http.StatusOK,
		Message: common.API_SUCCESS,
		Data:    map[string]interface{}{},
	}

	// Upvoting comment created in TestCreateComment()
	queryStr := "?" + "comment_id=" + commentID

	req, err := t_utils.MakeRequest(t_utils.DELETE, comments.COMMENTS_ROUTE_DELETE+queryStr, nil)
	if err != nil {
		return
	}

	// add the custom headers
	for key, value := range customHeaders {
		req.Header.Add(key, value)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		log.Println(err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected.Status, resp.Status)
	assert.Equal(t, expected.Message, resp.Message)

	mongoDeletionResult, ok := resp.Data["deleted"]
	if !ok {
		log.Fatal("DeleteComment does not have expected response body. Comment deletion has likely failed.")
	} else {
		count, ok := mongoDeletionResult.(map[string]interface{})["ModifiedCount"]
		if ok {
			assert.Equal(t, float64(1), count)
			log.Println("ModifiedCount == 1, which means the comment has been deleted")
		}
		log.Println("DeleteComment has expected response body. Comment deleted successfully.")
	}
}
