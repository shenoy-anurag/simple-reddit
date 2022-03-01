package communities

import (
	"context"
	"fmt"
	"net/http"
	"simple-reddit/configs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const COMMUNITY_ROUTE_PREFIX = "/community"

const CommunitiesCollectionName string = "communities"

var communityCollection *mongo.Collection = configs.GetCollection(configs.MongoClient, CommunitiesCollectionName)
var validate = validator.New()

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

func retrieveCommunityDetails(commReq GetCommunityRequest) (CommunityResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var community CommunityDBModel
	var commResp CommunityResponse
	filter := bson.D{primitive.E{Key: "name", Value: commReq.Name}}
	err := communityCollection.FindOne(ctx, filter).Decode(&community)
	if err != nil {
		return commResp, err
	}
	commResp, err = ConvertCommunityDBModelToCommunityResponse(community)
	return commResp, err
}

func checkCommunityNameExists(communityName string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var alreadyExists bool
	var community CommunityDBModel
	filter := bson.D{primitive.E{Key: "name", Value: communityName}}
	err := communityCollection.FindOne(ctx, filter).Decode(&community)
	if err == nil {
		if community.Name == communityName {
			alreadyExists = true
		} else {
			alreadyExists = false
		}
	} else {
		if err == mongo.ErrNoDocuments {
			err = nil
		}
	}
	return alreadyExists, err
}

func CreateCommunity() gin.HandlerFunc {
	return func(c *gin.Context) {
		var communityReq CreateCommunityRequest

		// validate the request body
		if err := c.BindJSON(&communityReq); err != nil {
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
		if validationErr := validate.Struct(&communityReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				configs.APIResponse{
					Status:  http.StatusBadRequest,
					Message: configs.API_FAILURE,
					Data:    map[string]interface{}{"data": validationErr.Error()}},
			)
			return
		}

		communityAlreadyExists, _ := checkCommunityNameExists(communityReq.Name)
		if communityAlreadyExists {
			c.JSON(
				http.StatusOK,
				configs.APIResponse{
					Status:  http.StatusOK,
					Message: configs.API_FAILURE,
					Data: map[string]interface{}{
						"message":                "Community with that name already exists",
						"communityAlreadyExists": communityAlreadyExists,
					}},
			)
			return
		}

		result, err := createCommunityInDB(communityReq)
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
				Data:    map[string]interface{}{"data": result, "community_name": communityReq.Name}},
		)
	}
}

func GetCommunity() gin.HandlerFunc {
	return func(c *gin.Context) {
		var communityReq GetCommunityRequest

		// validate the request body
		if err := c.BindUri(&communityReq); err != nil {
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
		if validationErr := validate.Struct(&communityReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				configs.APIResponse{
					Status:  http.StatusBadRequest,
					Message: configs.API_FAILURE,
					Data:    map[string]interface{}{"data": validationErr.Error()}},
			)
			return
		}
		fmt.Println(communityReq)
		communityDetails, err := retrieveCommunityDetails(communityReq)
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

		if communityDetails.ID != primitive.NilObjectID {
			c.JSON(
				http.StatusOK,
				configs.APIResponse{
					Status:  http.StatusOK,
					Message: configs.API_SUCCESS,
					Data:    map[string]interface{}{"community": communityDetails}},
			)
			return
		} else {
			c.JSON(
				http.StatusOK,
				configs.APIResponse{
					Status:  http.StatusOK,
					Message: configs.API_SUCCESS,
					Data:    map[string]interface{}{"community": communityDetails}},
			)
			return
		}
	}
}

func Routes(router *gin.Engine) {
	router.POST(COMMUNITY_ROUTE_PREFIX, CreateCommunity())
	router.GET(COMMUNITY_ROUTE_PREFIX+"/:name", GetCommunity())
}
