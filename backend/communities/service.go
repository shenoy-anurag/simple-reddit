package communities

import (
	"context"
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
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COMMUNITY_ROUTE_PREFIX = "/community"

const CommunitiesCollectionName string = "communities"
const PostsCollectionName string = "posts"

var CommunityCollection *mongo.Collection = configs.GetCollection(configs.MongoDB, CommunitiesCollectionName)
var PostsCollection *mongo.Collection = configs.GetCollection(configs.MongoDB, PostsCollectionName)
var validate = validator.New()

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
		// communityReq.IsUser = bool(c.Request.URL.Query().Get("isuser"))
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
		// if validationErr := validate.Struct(&communityReq); validationErr != nil {
		// 	c.JSON(
		// 		http.StatusBadRequest,
		// 		common.APIResponse{
		// 			Status:  http.StatusBadRequest,
		// 			Message: common.API_FAILURE,
		// 			Data:    map[string]interface{}{"error": validationErr.Error()}},
		// 	)
		// 	return
		// }
		//communityReq.Name = c.Request.URL.Query().Get("name")
		if communityReq.IsUser {
			allCommunities, err := retrieveAllCommunitiesOfUser(communityReq)
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
			if len(allCommunities) > 0 {
				c.JSON(
					http.StatusOK,
					common.APIResponse{
						Status:  http.StatusOK,
						Message: common.API_SUCCESS,
						Data:    map[string]interface{}{"communitites": allCommunities}},
				)
				return
			} else {
				c.JSON(
					http.StatusOK,
					common.APIResponse{
						Status:  http.StatusNotFound,
						Message: common.API_SUCCESS,
						Data:    map[string]interface{}{"communitites": allCommunities}},
				)
				return
			}
		}
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

func GetAllCommunities() gin.HandlerFunc {
	return func(c *gin.Context) {
		var getCommunityReq GetAllCommunitiesRequest
		// validate the request body
		if err := c.BindQuery(&getCommunityReq); err != nil {
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
		if validationErr := validate.Struct(&getCommunityReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}
		getCommunityReq.fill_defaults()

		communities, docCount, err := retrieveAllCommunitiesDetails(getCommunityReq)
		if err == mongo.ErrNoDocuments {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": common.ERR_COMMUNITY_NOT_FOUND.Message}},
			)
			return
		} else if err != nil {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusInternalServerError,
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
				Data:    map[string]interface{}{"num_communities": docCount, "communities": communities}},
		)
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
		if !communityAlreadyExists {
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

func GetCommunityPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var commPostReq GetPostsRequest

		// validate the request body
		if err := c.BindJSON(&commPostReq); err != nil {
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
		if validationErr := validate.Struct(&commPostReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}
		communityPosts, err := retrieveAllPosts(commPostReq)
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
		c.JSON(
			http.StatusOK,
			common.APIResponse{
				Status:  http.StatusOK,
				Message: common.API_SUCCESS,
				Data:    map[string]interface{}{"posts": communityPosts}},
		)
	}
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
				Message: common.API_SUCCESS,
				Data:    map[string]interface{}{"communityAlreadyExists": communityAlreadyExists},
			},
		)
	}
}

func createCommunityInDB(community CreateCommunityRequest) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newCommunity := ConvertCommunityRequestToCommunityDBModel(community) //, err := ConvertCommunityRequestToCommunityDBModel(community)
	// if err != nil {
	// 	return result, err
	// }
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

func retrieveAllCommunitiesDetails(getCommunityReq GetAllCommunitiesRequest) ([]CommunityResponse, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var communities []CommunityDBModel
	var communitiesResponses []CommunityResponse

	// Count the total number of documents in order to decide the number of pages.
	filter := bson.D{}
	docCount, err := CommunityCollection.CountDocuments(ctx, filter)
	if err != nil {
		return communitiesResponses, docCount, err
	}
	// Retrieve documents required for the page.
	var numItemsToSkip uint32 = (getCommunityReq.PageNumber - 1) * getCommunityReq.ItemsPerPage
	queryOptions := options.Find().SetSkip(int64(numItemsToSkip)).SetLimit(int64(getCommunityReq.ItemsPerPage))
	cursor, err := CommunityCollection.Find(ctx, filter, queryOptions)
	if err != nil {
		return communitiesResponses, docCount, err
	}
	if err = cursor.All(ctx, &communities); err != nil {
		return communitiesResponses, docCount, err
	}
	for _, community := range communities {
		item := ConvertCommunityDBModelToCommunityResponse(community)
		communitiesResponses = append(communitiesResponses, item)
	}
	return communitiesResponses, docCount, err
}

func retrieveAllCommunitiesOfUser(commReq GetCommunityRequest) ([]CommunityResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var communities []CommunityDBModel
	var communitiesResponses []CommunityResponse
	filter := bson.M{"username": commReq.Name}
	cursor, err := CommunityCollection.Find(ctx, filter)
	if err != nil {
		return communitiesResponses, err
	}
	if err = cursor.All(ctx, &communities); err != nil {
		return communitiesResponses, err
	}
	for _, community := range communities {
		item := ConvertCommunityDBModelToCommunityResponse(community)
		if err != nil {
			return communitiesResponses, err
		}
		communitiesResponses = append(communitiesResponses, item)
	}
	return communitiesResponses, err
}

func retrieveAllPosts(postReq GetPostsRequest) ([]PostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var posts []PostDBModel
	var postResp []PostResponse
	var community CommunityDBModel
	communityFilter := bson.D{primitive.E{Key: "name", Value: postReq.Name}}
	err := CommunityCollection.FindOne(ctx, communityFilter).Decode(&community)
	if err != nil {
		return postResp, err
	}
	postFilter := bson.M{"community_id": community.ID}
	cursor, err := PostsCollection.Find(ctx, postFilter)
	if err != nil {
		return postResp, err
	}
	if err = cursor.All(ctx, &posts); err != nil {
		return postResp, err
	}
	for _, post := range posts {
		item, err := ConvertPostDBModelToPostResponse(post)
		if err != nil {
			return postResp, err
		}
		postResp = append(postResp, item)
	}
	return postResp, err
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

func Routes(router *gin.Engine) {
	router.POST(COMMUNITY_ROUTE_PREFIX+"/create", CreateCommunity())
	router.POST(COMMUNITY_ROUTE_PREFIX+"/check-name", CheckCommunityExists())
	// router.GET(COMMUNITY_ROUTE_PREFIX, GetCommunity())
	router.POST(COMMUNITY_ROUTE_PREFIX, GetCommunity())
	// TODO - user who have following a community - POST - username
	router.POST(COMMUNITY_ROUTE_PREFIX+"/all", GetAllCommunities())  //all community that exists
	// router.GET(COMMUNITY_ROUTE_PREFIX+"/home", GetCommunityPosts())
	//router.POST(COMMUNITY_ROUTE_PREFIX+"/home", GetCommunityPosts())
	router.PATCH(COMMUNITY_ROUTE_PREFIX, EditCommunity())
	// router.DELETE(COMMUNITY_ROUTE_PREFIX, DeleteCommunity())
	router.POST(COMMUNITY_ROUTE_PREFIX+"/delete", DeleteCommunity())
}
