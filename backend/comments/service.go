package comments

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

const COMMENTS_ROUTE_PREFIX = "/comment"

const CommentsCollectionName string = "comments"

var CommentsCollection *mongo.Collection = configs.GetCollection(configs.MongoDB, CommentsCollectionName)
var validate = validator.New()

func CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var commentReq CreateCommentRequest

		// validate the request body
		if err := c.BindJSON(&commentReq); err != nil {
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
		if validationErr := validate.Struct(&commentReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}

		result, err := createCommentInDB(commentReq)
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
				Data:    map[string]interface{}{"created": result}},
		)
	}
}

func createCommentInDB(comment CreateCommentRequest) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newComment, err := ConvertCommentRequestToCommentDBModel(comment) //, err := ConvertCommunityRequestToCommunityDBModel(community)
	if err != nil {
		return result, err
	}
	result, err = CommentsCollection.InsertOne(ctx, newComment)
	return result, err
}

func Routes(router *gin.Engine) {
	router.POST(COMMENTS_ROUTE_PREFIX, CreateComment())
}
