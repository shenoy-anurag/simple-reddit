package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"simple-reddit/common"
	"simple-reddit/profiles"
	t_utils "simple-reddit/test_utils"
	"simple-reddit/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	var resp = common.APIResponse{}
	var body = users.CreateUserRequest{
		FirstName: "Mary",
		LastName:  "Jane",
		Email:     "mary.jane@ufl.edu",
		Username:  "mary_jane",
		Password:  "password123",
	}
	var expected = common.APIResponse{
		Status:  http.StatusCreated,
		Message: common.API_SUCCESS,
		Data:    map[string]interface{}{},
	}

	req, err := t_utils.MakeRequest(t_utils.POST, users.USER_ROUTE_PREFIX+"/signup", body)
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

	log.Println(resp)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, expected.Status, resp.Status)
	assert.Equal(t, expected.Message, resp.Message)
	assert.Contains(t, resp.Data, "data")
}

func TestLogin(t *testing.T) {
	var resp = common.APIResponse{}
	var body = users.LoginUserRequest{
		Username: "mary_jane",
		Password: "password123",
	}
	userDB, _ := users.GetUserDetails(body.Username)
	userResp := users.ConvertUserDBModelToUserResponse(userDB)
	var expected = common.APIResponse{
		Status:  http.StatusOK,
		Message: common.API_SUCCESS,
		Data:    map[string]interface{}{"accessToken": "", "user": userResp},
	}

	// // isoDateTime, err := time.Parse("YYYY-mm-ddTHH:MM:ssZ", "2022-01-21T04:41:56.616Z")
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

	req, err := t_utils.MakeRequest(t_utils.POST, users.USER_ROUTE_PREFIX+"/loginuser", body)
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
	assert.Equal(t, expected.Data["user"], userResp)
}

func TestDeleteProfileForUser(t *testing.T) {
	var resp = common.APIResponse{}
	var body = profiles.DeleteProfileRequest{
		Username: "mary_jane",
		Password: "password123",
	}
	var expected = common.APIResponse{
		Status:  http.StatusOK,
		Message: common.API_SUCCESS,
		Data:    map[string]interface{}{},
	}

	req, err := t_utils.MakeRequest(t_utils.POST, profiles.PROFILE_ROUTE_PREFIX+"/delete", body)
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

	log.Println(resp)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected.Status, resp.Status)
	assert.Equal(t, expected.Message, resp.Message)
	assert.Contains(t, resp.Data, "deletedProfile")
	assert.Contains(t, resp.Data, "deletedUser")
}
