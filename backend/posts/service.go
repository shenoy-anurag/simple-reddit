package posts

import (
	"context"
	"net/http"
	"simple-reddit/common"
	"simple-reddit/configs"
	"simple-reddit/users"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const POST_ROUTE_PREFIX = "/post"

const PostsCollectionName string = "posts"

var postCollection *mongo.Collection = configs.GetCollection(configs.MongoClient, PostsCollectionName)
var validate = validator.New()

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post CreatePostRequest
		if err := c.BindJSON(&post); err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}
		if validationErr := validate.Struct(&post); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}
		result, err := createPostInDB(post)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				common.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}

		c.JSON(
			http.StatusCreated,
			common.APIResponse{
				Status:  http.StatusCreated,
				Message: common.API_SUCCESS,
				Data:    map[string]interface{}{"error": result}},
		)

	}
}

func DeleteCommunity() gin.HandlerFunc {
	return func(c *gin.Context) {
		var delPostReq DeletePostRequest

		// validate the request body
		if err := c.BindJSON(&delPostReq); err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&delPostReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}

		// TODO: add check whether correct user is deleting the community after adding JWT verification.
		user, err := users.GetUserDetails(delPostReq.UserName)
		// TODO: replace this check with a check against username within claims of JWT token.
		if user.Username != delPostReq.UserName {
			c.JSON(
				http.StatusUnauthorized,
				common.APIResponse{
					Status:  http.StatusUnauthorized,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}
		result, err := deletePost(delPostReq)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				common.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			common.APIResponse{
				Status:  http.StatusOK,
				Message: common.API_SUCCESS,
				Data:    map[string]interface{}{"deleted": result}},
		)
	}
}

func GetPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postReq GetPostRequest
		// validate the request body
		if err := c.BindJSON(&postReq); err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}
		// use the validator library to validate required fields
		if validationErr := validate.Struct(&postReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}
		postDetails, err := retrievePostDetails(postReq)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				common.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}

		if len(postDetails) > 0 {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_SUCCESS,
					Data:    map[string]interface{}{"posts": postDetails}},
			)
			return
		} else {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusNotFound,
					Message: common.API_SUCCESS,
					Data:    map[string]interface{}{"posts": postDetails}},
			)
			return
		}
	}
}

func createPostInDB(post CreatePostRequest) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newPost := ConvertPostRequestToPostDBModel(post)
	if err != nil {
		return result, err
	}
	result, err = postCollection.InsertOne(ctx, newPost)
	return result, err
}

func retrievePostDetails(postReq GetPostRequest) ([]PostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var posts []PostDBModel
	var postResp []PostResponse
	filter := bson.M{"$and": []bson.M{{"username": postReq.UserName}, {"community_id": postReq.CommunityID}}} //bson.D{primitive.E{Key: "community_id", Value: postReq.CommunityID}, primitive.E{Key: "username", Value: postReq.UserName}}
	cursor, err := postCollection.Find(ctx, filter)
	if err = cursor.All(ctx, &posts); err != nil {
		return postResp, err
	}
	for _, post := range posts {
		item, err := ConvertPostDBModelToPostResponse(post)
		if err != nil {
			return postResp, err
		}
		postResp = append(postResp, item)
	}
	return postResp, err
}

func deletePost(postReq DeletePostRequest) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "username", Value: postReq.ID}}
	result, err := postCollection.DeleteOne(ctx, filter)
	return result, err
}

func Routes(router *gin.Engine) {
	router.POST(POST_ROUTE_PREFIX+"/create", CreatePost())
	router.GET(POST_ROUTE_PREFIX+"/get", GetPosts())
}
