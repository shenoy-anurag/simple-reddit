package posts

import (
	"context"
	"net/http"
	"simple-reddit/common"
	"simple-reddit/configs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

const POST_ROUTE_PREFIX = "/posts"

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
					Data:    map[string]interface{}{"data": err.Error()}},
			)
			return
		}
		if validationErr := validate.Struct(&post); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"data": validationErr.Error()}},
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
					Data:    map[string]interface{}{"data": err.Error()}},
			)
			return
		}

		c.JSON(
			http.StatusCreated,
			common.APIResponse{
				Status:  http.StatusCreated,
				Message: common.API_SUCCESS,
				Data:    map[string]interface{}{"data": result}},
		)

	}
}

func createPostInDB(post CreatePostRequest) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newPost, err := ConvertPostRequestToPostDBModel(post)
	if err != nil {
		return result, err
	}
	result, err = postCollection.InsertOne(ctx, newPost)
	return result, err
}

func Routes(router *gin.Engine) {
	router.POST(POST_ROUTE_PREFIX+"/create", CreatePost())
}
