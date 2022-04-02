package profiles

import (
	"context"
	"net/http"
	"simple-reddit/common"
	"simple-reddit/configs"
	// "simple-reddit/comments"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const PROFILE_ROUTE_PREFIX = "/profile"

const ProfilesCollectionName string = "profiles"
const UsersCollectionName string = "users"
const CommentsCollectionName string = "comments"
const SavedCollectionName string ="saved"
const PostsCollectionName string = "posts"

var ProfileCollection *mongo.Collection = configs.GetCollection(configs.MongoDB, ProfilesCollectionName)
var UsersCollection *mongo.Collection = configs.GetCollection(configs.MongoDB, UsersCollectionName)
var CommentsCollection *mongo.Collection = configs.GetCollection(configs.MongoDB, CommentsCollectionName)
var SavedCollection *mongo.Collection = configs.GetCollection(configs.MongoDB,SavedCollectionName)
var PostsCollection *mongo.Collection = configs.GetCollection(configs.MongoDB, PostsCollectionName)
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

func DeleteProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var delProfileReq DeleteProfileRequest

		// validate the request body
		if err := c.BindJSON(&delProfileReq); err != nil {
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
		if validationErr := validate.Struct(&delProfileReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}
		// update posts that the user has created
		postsResult, verifiedPosts, err:=updatePostsForDeletedUser(delProfileReq)
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
		if !verifiedPosts {
			c.JSON(
				http.StatusInternalServerError,
				common.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}
		// Delete the saved record for the user
		savedResult, verifiedSaved, err:= deleteSaved(delProfileReq)
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
		if !verifiedSaved {
			c.JSON(
				http.StatusInternalServerError,
				common.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}

			// Delete the profile record for the user
		profileResult, verifiedProfile, err:= deleteProfile(delProfileReq)
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
		if !verifiedProfile {
			c.JSON(
				http.StatusInternalServerError,
				common.APIResponse{
					Status:  http.StatusInternalServerError,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}

			// Delete the user record for the user
		userResult,verifiedUser, err := deleteUser(delProfileReq)
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
		if !verifiedUser {
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
				Data:    map[string]interface{}{"deletedProfile": profileResult,"deletedUser":userResult,"deletedSaved":savedResult,"updated_posts":postsResult, "username":delProfileReq.Username }},
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

func deleteProfile(delProfileReq DeleteProfileRequest) (*mongo.DeleteResult,bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "username", Value: delProfileReq.Username}}
	result, err := ProfileCollection.DeleteOne(ctx, filter)
	if err!=nil {
		return result,false,err
	}
	return result,true,err
}

func deleteUser(delProfileReq DeleteProfileRequest) (userResult *mongo.DeleteResult,status bool,err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "username", Value: delProfileReq.Username}}
	userDB, err := getUserDetails(delProfileReq.Username)
	if err!=nil {
		return userResult,false, err
	}
	userResult,verifiedUser,err := login(delProfileReq,userDB)
	if !verifiedUser {
		return userResult,false,err
	}
	userResult, err = UsersCollection.DeleteOne(ctx, filter)
	return userResult,true, err
}

func deleteSaved(delProfileReq DeleteProfileRequest) (*mongo.DeleteResult,bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "username", Value: delProfileReq.Username}}
	result, err := SavedCollection.DeleteOne(ctx, filter)
	if err!=nil {
		return result,false,err
	}
	return result,true,err
}

func updatePostsForDeletedUser(delProfileReq DeleteProfileRequest) (*mongo.UpdateResult,bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "username", Value: delProfileReq.Username}}
	updateQuery := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "username", Value: "deleted-user"},
				// primitive.E{Key: "lastname", Value: editProfileReq.LastName},
				// primitive.E{Key: "email", Value: editProfileReq.Email},
			},
		},
	}
	result, err := PostsCollection.UpdateMany(ctx, filter, updateQuery)
	// result, err = ProfileCollection.UpdateOne(ctx, filter, updateQuery)
	if err!=nil {
		return result,false,err
	}
	return result,true,err
}

func getUserDetails(userName string) (UserDBModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user UserDBModel
	filter := bson.D{primitive.E{Key: "username", Value: userName}}
	err := UsersCollection.FindOne(ctx, filter).Decode(&user)
	return user, err
}

func login(profileReq DeleteProfileRequest,userDB UserDBModel) (resultUser *mongo.DeleteResult,status bool, err error){
	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(profileReq.Password))
	if err != nil {
		return resultUser, false,err
	}
	return resultUser,true,err
}

func CreateSavedPC(UserName string)(SavedDBModel){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	SavedDB := CreateSavedDBModel(UserName)
	// if err != nil {
	// 	return result, err
	// }
	_, _ = SavedCollection.InsertOne(ctx, SavedDB)
	return SavedDB

}

func Routes(router *gin.Engine) {
	router.POST(PROFILE_ROUTE_PREFIX, GetProfile())
	router.PATCH(PROFILE_ROUTE_PREFIX, EditProfile()) // maybe PATCH > DELETE
	router.POST(PROFILE_ROUTE_PREFIX+"/delete", DeleteProfile()) // router.DELETE(PROFILE_ROUTE_PREFIX, DeleteProfile()) // DELETE -> POST
	// router.PATCH(PROFILE_ROUTE_PREFIX+"/comment", SaveComment()) // maybe PATCH > DELETE
}
