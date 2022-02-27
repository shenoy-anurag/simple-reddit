package communities

import (
	"context"
	"net/http"
	"simple-reddit/configs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

const COMMUNITY_ROUTE_PREFIX = "/community"

const CommunitiesCollectionName string = "communities"

var communityCollection *mongo.Collection = configs.GetCollection(configs.MongoClient, CommunitiesCollectionName)
var validate = validator.New()

func CreateCommunity() gin.HandlerFunc {
	return func(c *gin.Context) {
		var community CreateCommunityRequest

		// validate the request body
		if err := c.BindJSON(&community); err != nil {
			c.JSON(
				http.StatusBadRequest,
				configs.APIResponse{
					Status:  http.StatusBadRequest,
					Message: configs.API_FAILURE,
					Data:    map[string]interface{}{"data": err.Error()}},
			)
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&community); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				configs.APIResponse{
					Status:  http.StatusBadRequest,
					Message: configs.API_FAILURE,
					Data:    map[string]interface{}{"data": validationErr.Error()}},
			)
			return
		}

		result, err := createCommunityInDB(community)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				configs.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: configs.API_ERROR,
					Data:    map[string]interface{}{"data": err.Error()}},
			)
			return
		}

		c.JSON(
			http.StatusCreated,
			configs.APIResponse{
				Status:  http.StatusCreated,
				Message: configs.API_SUCCESS,
				Data:    map[string]interface{}{"data": result, "community_name": community.Name}},
		)
	}
}

func createCommunityInDB(community CreateCommunityRequest) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newCommunity, err := ConvertCommunityRequestToCommunityDBModel(community)
	if err != nil {
		return result, err
	}
	result, err = communityCollection.InsertOne(ctx, newCommunity)
	return result, err
}

func Routes(router *gin.Engine) {
	router.POST(COMMUNITY_ROUTE_PREFIX, CreateCommunity())
}
