package communities

import (
	"context"
	"fmt"
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

const COMMUNITY_ROUTE_PREFIX = "/community"

const CommunitiesCollectionName string = "communities"

var CommunityCollection *mongo.Collection = configs.GetCollection(configs.MongoDB, CommunitiesCollectionName)
var validate = validator.New()

func createCommunityInDB(community CreateCommunityRequest) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	newCommunity, err := ConvertCommunityRequestToCommunityDBModel(community)
	if err != nil {
		return result, err
	}
	result, err = CommunityCollection.InsertOne(ctx, newCommunity)
	return result, err
}

func retrieveCommunityDetails(commReq GetCommunityRequest) (CommunityDBModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var community CommunityDBModel
	filter := bson.D{primitive.E{Key: "name", Value: commReq.Name}}
	err := CommunityCollection.FindOne(ctx, filter).Decode(&community)
	return community, err
}

func editCommunityDetails(communityReq EditCommunityRequest) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	getCommunityReq := ConvertEditCommunityReqToGetCommunityReq(communityReq)

	community, err := retrieveCommunityDetails(getCommunityReq)
	if err != nil {
		return result, err
	}
	// updating the data in db
	filter := bson.M{"_id": community.ID}
	updateQuery := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "description", Value: communityReq.Description},
			},
		},
	}
	result, err = CommunityCollection.UpdateOne(
		ctx,
		filter,
		updateQuery,
	)
	return result, err
}

func deleteCommunity(commReq DeleteCommunityRequest) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "name", Value: commReq.Name}}
	result, err := CommunityCollection.DeleteOne(ctx, filter)
	return result, err
}

func checkCommunityNameExists(communityName string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var alreadyExists bool
	var community CommunityDBModel
	filter := bson.D{primitive.E{Key: "name", Value: communityName}}
	err := CommunityCollection.FindOne(ctx, filter).Decode(&community)
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

func CheckCommunityExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		var communityReq GetCommunityRequest

		// validate the request body
		if err := c.BindJSON(&communityReq); err != nil {
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
		if validationErr := validate.Struct(&communityReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}

		communityAlreadyExists, err := checkCommunityNameExists(communityReq.Name)
		if err != nil {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_ERROR,
					Data: map[string]interface{}{
						"error": err.Error(),
					},
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			common.APIResponse{
				Status:  http.StatusOK,
				Message: common.API_FAILURE,
				Data:    map[string]interface{}{"communityAlreadyExists": communityAlreadyExists},
			},
		)
	}
}

func CreateCommunity() gin.HandlerFunc {
	return func(c *gin.Context) {
		var communityReq CreateCommunityRequest

		// validate the request body
		if err := c.BindJSON(&communityReq); err != nil {
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
		if validationErr := validate.Struct(&communityReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}

		communityAlreadyExists, err := checkCommunityNameExists(communityReq.Name)
		if communityAlreadyExists {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_FAILURE,
					Data: map[string]interface{}{
						"error":                  common.ERR_COMMUNITY_ALREADY_EXISTS.Message,
						"communityAlreadyExists": communityAlreadyExists,
					}},
			)
			return
		} else if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				common.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}

		result, err := createCommunityInDB(communityReq)
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
				Data:    map[string]interface{}{"created": result, "community_name": communityReq.Name}},
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
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&communityReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}
		fmt.Println(communityReq)
		communityDB, err := retrieveCommunityDetails(communityReq)
		if err == mongo.ErrNoDocuments {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": common.ERR_COMMUNITY_NOT_FOUND.Message}},
			)
			return
		}
		communityDetails := ConvertCommunityDBModelToCommunityResponse(communityDB)
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

		if communityDetails.ID != primitive.NilObjectID {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_SUCCESS,
					Data:    map[string]interface{}{"community": communityDetails}},
			)
			return
		} else {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_SUCCESS,
					Data:    map[string]interface{}{"community": communityDetails}},
			)
			return
		}
	}
}

func EditCommunity() gin.HandlerFunc {
	return func(c *gin.Context) {
		var communityReq EditCommunityRequest

		// validate the request body
		if err := c.BindJSON(&communityReq); err != nil {
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
		if validationErr := validate.Struct(&communityReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}

		communityAlreadyExists, err := checkCommunityNameExists(communityReq.Name)
		if communityAlreadyExists {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_FAILURE,
					Data: map[string]interface{}{
						"message":                "Community with that name already exists",
						"communityAlreadyExists": communityAlreadyExists,
					}},
			)
			return
		} else if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				common.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}

		result, err := editCommunityDetails(communityReq)
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
				Data:    map[string]interface{}{"updated": result, "community_name": communityReq.Name}},
		)
	}
}

func DeleteCommunity() gin.HandlerFunc {
	return func(c *gin.Context) {
		var delCommunityReq DeleteCommunityRequest

		// validate the request body
		if err := c.BindJSON(&delCommunityReq); err != nil {
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
		if validationErr := validate.Struct(&delCommunityReq); validationErr != nil {
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
		user, err := users.GetUserDetails(delCommunityReq.UserName)
		// TODO: replace this check with a check against username within claims of JWT token.
		if user.Username != delCommunityReq.UserName {
			c.JSON(
				http.StatusUnauthorized,
				common.APIResponse{
					Status:  http.StatusUnauthorized,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}
		result, err := deleteCommunity(delCommunityReq)
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
				Data:    map[string]interface{}{"deleted": result, "community_name": delCommunityReq.Name}},
		)
	}
}

func Routes(router *gin.Engine) {
	router.POST(COMMUNITY_ROUTE_PREFIX, CreateCommunity())
	router.POST(COMMUNITY_ROUTE_PREFIX+"/check-name", CheckCommunityExists())
	router.GET(COMMUNITY_ROUTE_PREFIX+"/:name", GetCommunity())
	router.PATCH(COMMUNITY_ROUTE_PREFIX, EditCommunity())
	router.DELETE(COMMUNITY_ROUTE_PREFIX, DeleteCommunity())
}
