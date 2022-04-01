package comments

import (
	"net/http"

	"simple-reddit/common"
	"simple-reddit/configs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

const COMMENTS_ROUTE_PREFIX = "/comment"

const CommentsCollectionName string = "comments"
const CommentsVotingHistoryCollectionName string = "comments_voting_history"

var CommentsCollection *mongo.Collection = configs.GetCollection(configs.MongoDB, CommentsCollectionName)
var CommentsVotingHistoryCollection *mongo.Collection = configs.GetCollection(configs.MongoDB, CommentsVotingHistoryCollectionName)
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

		// Check whether Parent Comment exists or is deleted.
		if commentReq.ParentId != "" {
			parentComment, err := retrieveCommentById(commentReq.ParentId)
			if err != nil {
				c.JSON(
					http.StatusNotFound,
					common.APIResponse{
						Status:  http.StatusNotFound,
						Message: common.API_FAILURE,
						Data:    map[string]interface{}{"error": err.Error(), "message": "parent comment not found"}},
				)
				return
			} else if parentComment.IsDeleted {
				c.JSON(
					http.StatusOK,
					common.APIResponse{
						Status:  http.StatusOK,
						Message: common.API_FAILURE,
						Data:    map[string]interface{}{"error": common.ERR_PARENT_COMMENT_IS_DELETED.Message}},
				)
				return
			}
		}

		// Create comment in database
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

func DeleteComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var delCommentReq DeleteCommentRequest

		// validate the request body
		if err := c.BindJSON(&delCommentReq); err != nil {
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
		if validationErr := validate.Struct(&delCommentReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}

		result, err := deleteComment(delCommentReq)
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

func VoteComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cVoteReq CommentVoteRequest

		// validate the request body
		if err := c.BindJSON(&cVoteReq); err != nil {
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
		if validationErr := validate.Struct(&cVoteReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}

		// Update Vote in DB
		result, err := updateVote(cVoteReq)
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

func Routes(router *gin.Engine) {
	router.POST(COMMENTS_ROUTE_PREFIX, CreateComment())
	router.POST(COMMENTS_ROUTE_PREFIX+"/vote", CreateComment())
	router.DELETE(COMMENTS_ROUTE_PREFIX, DeleteComment())
}
