package users

import (
	"context"
	"net/http"
	"simple-reddit/configs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

const UsersCollectionName string = "users"

var userCollection *mongo.Collection = configs.GetCollection(configs.MongoClient, UsersCollectionName)
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user CreateUserRequest
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, configs.APIResponse{Status: http.StatusBadRequest, Message: configs.API_ERROR, Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, configs.APIResponse{Status: http.StatusBadRequest, Message: configs.API_ERROR, Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newUserStruct := ConvertUserRequestToUserDBModel(user)

		result, err := userCollection.InsertOne(ctx, newUserStruct)
		if err != nil {
			c.JSON(http.StatusInternalServerError, configs.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, configs.APIResponse{Status: http.StatusCreated, Message: configs.API_SUCCESS, Data: map[string]interface{}{"data": result}})
	}
}

func Routes(router *gin.Engine) {
	router.POST("/signup", CreateUser())
}
