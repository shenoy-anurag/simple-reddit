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

func TestSignupForProfile(t *testing.T) {
	var resp = common.APIResponse{}
	var body = users.CreateUserRequest{
		FirstName: "Albert",
		LastName:  "Einstein",
		Email:     "einsteinalbert@ufl.edu",
		Username:  "albert",
		Password:  "special-relativity",
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

func TestGetProfile(t *testing.T) {
	var resp = common.APIResponse{}
	var body = profiles.GetProfileRequest{
		UserName: "albert",
	}
	var expected = common.APIResponse{
		Status:  http.StatusOK,
		Message: common.API_SUCCESS,
		Data:    map[string]interface{}{},
	}

	req, err := t_utils.MakeRequest(t_utils.POST, profiles.PROFILE_ROUTE_PREFIX, body)
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

	log.Println("API Response", resp)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected.Status, resp.Status)
	assert.Equal(t, expected.Message, resp.Message)
	assert.Contains(t, resp.Data, "Profile")
}

func TestEditProfile(t *testing.T) {
	var resp = common.APIResponse{}
	var body = profiles.EditProfileRequest{
		FirstName: "Albert",
		LastName:  "Einstein",
		Email:     "einsteinalbert@ufl.edu",
		UserName:  "albert",
	}
	var expected = common.APIResponse{
		Status:  http.StatusOK,
		Message: common.API_SUCCESS,
		Data:    map[string]interface{}{},
	}

	req, err := t_utils.MakeRequest(t_utils.PATCH, profiles.PROFILE_ROUTE_PREFIX, body)
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
	assert.Contains(t, resp.Data, "updated")
}

func TestDeleteProfile(t *testing.T) {
	var resp = common.APIResponse{}
	var body = profiles.DeleteProfileRequest{
		Username: "albert",
		Password: "special-relativity",
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
