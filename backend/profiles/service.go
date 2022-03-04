package profiles

import (
	"simple-reddit/configs"
	"context"
	"time"
	
	"github.com/go-playground/validator/v10"
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

func CreateProfileInDB(profileReq ProfileDBModel) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err = ProfileCollection.InsertOne(ctx, profileReq)
	return result, err
}
