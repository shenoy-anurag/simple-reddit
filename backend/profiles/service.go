package profiles

import (
	"context"
	"net/http"
	"simple-reddit/common"
	"simple-reddit/configs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const PROFILE_ROUTE_PREFIX = "/profile"

const ProfilesCollectionName string = "profiles"

var ProfileCollection *mongo.Collection = configs.GetCollection(configs.MongoDB, ProfilesCollectionName)
var validate = validator.New()

func CreateProfile(profileReq ProfileDBModel) bool {
	_, err := CreateProfileInDB(profileReq)
	if err != nil {
		return false
	}
	return true
}

func GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var profileReq GetProfileRequest
		// validate the request body
		if err := c.BindJSON(&profileReq); err != nil {
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
		if validationErr := validate.Struct(&profileReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}
		profileDetails, err := retrieveProfileDetails(profileReq)
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
				Data:    map[string]interface{}{"Profile": profileDetails}},
		)
		return
	}
}

func EditProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var editprofiletReq EditProfileRequest

		// validate the request body
		if err := c.BindJSON(&editprofiletReq); err != nil {
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
		if validationErr := validate.Struct(&editprofiletReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}
		result, err := editProfileDetails(editprofiletReq)
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
				Data:    map[string]interface{}{"updated": result}},
		)
	}
}

func CreateProfileInDB(profileReq ProfileDBModel) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err = ProfileCollection.InsertOne(ctx, profileReq)
	return result, err
}

func retrieveProfileDetails(profileReq GetProfileRequest) (ProfileResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var profileDB ProfileDBModel
	filter := bson.M{"username": profileReq.UserName}
	err := ProfileCollection.FindOne(ctx, filter).Decode(&profileDB)
	profileResp := ConvertProfileDBModelToProfileResponse(profileDB)
	return profileResp, err
}

func editProfileDetails(editProfileReq EditProfileRequest) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// updating the data in db
	profileExists, err := checkProfileExists(editProfileReq.UserName)
	if !profileExists {
		return result, err
	}
	filter := bson.M{"username": editProfileReq.UserName}
	updateQuery := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "firstname", Value: editProfileReq.FirstName},
				primitive.E{Key: "lastname", Value: editProfileReq.LastName},
				primitive.E{Key: "email", Value: editProfileReq.Email},
			},
		},
	}
	result, err = ProfileCollection.UpdateOne(ctx, filter, updateQuery)
	return result, err
}

func checkProfileExists(UserNameReq string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var alreadyExists bool
	var profile ProfileDBModel
	filter := bson.D{primitive.E{Key: "username", Value: UserNameReq}}
	err := ProfileCollection.FindOne(ctx, filter).Decode(&profile)
	if err == nil {
		if profile.UserName == UserNameReq {
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

func Routes(router *gin.Engine) {
	router.POST(PROFILE_ROUTE_PREFIX, GetProfile())
	router.PATCH(PROFILE_ROUTE_PREFIX, EditProfile())
}
