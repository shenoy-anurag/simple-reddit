package users

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
	"golang.org/x/crypto/bcrypt"
)

const HASHING_COST = 10
const USER_ROUTE_PREFIX = "/users"

const UsersCollectionName string = "users"

var UsersCollection *mongo.Collection = configs.GetCollection(configs.MongoDB, UsersCollectionName)
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user CreateUserRequest

		// validate the request body
		if err := c.BindJSON(&user); err != nil {
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
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}

		saltedAndHashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), HASHING_COST)
		if err != nil {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}

		user.Password = string(saltedAndHashedPwd)
		usernameAlreadyExists, err := CheckUsername(user.Username)
		if usernameAlreadyExists {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_FAILURE,
					Data: map[string]interface{}{
						"error":                 common.ERR_USERNAME_ALREADY_EXISTS.Message,
						"usernameAlreadyExists": usernameAlreadyExists},
				},
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
		result, err := CreateUserInDB(user)
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
				Data:    map[string]interface{}{"data": result}},
		)
	}
}

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//var userName string
		//var userPwd string
		var loginUserReq LoginUserRequest
		if err := c.BindJSON(&loginUserReq); err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}
		userDB, err := GetUserDetails(loginUserReq.Username)
		if err != nil {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(loginUserReq.Password))
		if err != nil {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": common.ERR_INCORRECT_CREDENTIALS.Message}},
			)
			return
		}
		if err == nil {
			token := configs.JWTAuthService().GenerateToken(userDB.Username)

			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_SUCCESS,
					Data: map[string]interface{}{
						"accessToken": token,
						"user":        ConvertUserDBModelToUserResponse(userDB)},
				},
			)
		}
	}
}

func CreateUserInDB(user CreateUserRequest) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newUserStruct := ConvertUserRequestToUserDBModel(user)
	result, err = UsersCollection.InsertOne(ctx, newUserStruct)
	return result, err
}

// Provide username and context as parameter to
func GetUserDetails(userName string) (UserDBModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user UserDBModel
	filter := bson.D{primitive.E{Key: "username", Value: userName}}
	err := UsersCollection.FindOne(ctx, filter).Decode(&user)
	return user, err
}

func CheckUsername(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var alreadyExists bool
	var user UserDBModel
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	err := UsersCollection.FindOne(ctx, filter).Decode(&user)
	if err == nil {
		if user.Username == username {
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

func CheckUsernameExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user CheckUsernameRequest
		// validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}

		usernameAlreadyExists, err := CheckUsername(user.Username)
		if err != nil {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}
		c.JSON(
			http.StatusOK,
			common.APIResponse{
				Status:  http.StatusOK,
				Message: common.API_SUCCESS,
				Data:    map[string]interface{}{"usernameAlreadyExists": usernameAlreadyExists}},
		)
	}
}

func Routes(router *gin.Engine) {
	router.POST(USER_ROUTE_PREFIX+"/signup", CreateUser())
	router.POST(USER_ROUTE_PREFIX+"/loginuser", LoginUser())
	router.POST(USER_ROUTE_PREFIX+"/check-username", CheckUsernameExists())
}
