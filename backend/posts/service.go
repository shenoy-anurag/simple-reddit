package posts

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

const POST_ROUTE_PREFIX = "/post"
const HOME_ROUTE_PREFIX = "/home"

const PostsCollectionName string = "posts"

var PostsCollection *mongo.Collection = configs.GetCollection(configs.MongoDB, PostsCollectionName)
var validate = validator.New()

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post CreatePostRequest
		if err := c.BindJSON(&post); err != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}
		if validationErr := validate.Struct(&post); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}
		result, err := createPostInDB(post)
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

func DeletePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var delPostReq DeletePostRequest

		// validate the request body
		if err := c.BindJSON(&delPostReq); err != nil {
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
		if validationErr := validate.Struct(&delPostReq); validationErr != nil {
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
		user, err := users.GetUserDetails(delPostReq.UserName)
		// TODO: replace this check with a check against username within claims of JWT token.
		if user.Username != delPostReq.UserName {
			c.JSON(
				http.StatusUnauthorized,
				common.APIResponse{
					Status:  http.StatusUnauthorized,
					Message: common.API_ERROR,
					Data:    map[string]interface{}{"error": err.Error()}},
			)
			return
		}
		result, err := deletePost(delPostReq)
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

func GetPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postReq GetPostRequest
		// validate the request body
		if err := c.BindJSON(&postReq); err != nil {
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
		if validationErr := validate.Struct(&postReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}
		postDetails, err := retrievePostDetails(postReq)
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

		if len(postDetails) > 0 {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_SUCCESS,
					Data:    map[string]interface{}{"posts": postDetails}},
			)
			return
		} else {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusNotFound,
					Message: common.API_SUCCESS,
					Data:    map[string]interface{}{"posts": postDetails}},
			)
			return
		}
	}
}

func GetFeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var feedReq GetFeedRequest

		// validate the request body
		if err := c.BindJSON(&feedReq); err != nil {
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
		if validationErr := validate.Struct(&feedReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}

		postDetails, err := retrieveFeedDetails(feedReq) // retrieveAllPostDetails()
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

		if len(postDetails) > 0 {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusOK,
					Message: common.API_SUCCESS,
					Data:    map[string]interface{}{"posts": postDetails}},
			)
			return
		} else {
			c.JSON(
				http.StatusOK,
				common.APIResponse{
					Status:  http.StatusNotFound,
					Message: common.API_SUCCESS,
					Data:    map[string]interface{}{"posts": postDetails}},
			)
			return
		}
	}
}

func EditPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postReq EditPostRequest

		// validate the request body
		if err := c.BindJSON(&postReq); err != nil {
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
		if validationErr := validate.Struct(&postReq); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				common.APIResponse{
					Status:  http.StatusBadRequest,
					Message: common.API_FAILURE,
					Data:    map[string]interface{}{"error": validationErr.Error()}},
			)
			return
		}
		result, err := editPostDetails(postReq)
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

func createPostInDB(post CreatePostRequest) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newPost := ConvertPostRequestToPostDBModel(post)
	if err != nil {
		return result, err
	}
	result, err = PostsCollection.InsertOne(ctx, newPost)
	return result, err
}

func CheckPostExists(postReq DeletePostRequest) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var post PostDBModel
	filter := bson.M{"$and": []bson.M{{"username": postReq.UserName}, {"_id": postReq.ID}}}
	//cursor, err := postCollection.FindOne(ctx, filter)
	err := PostsCollection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		return false, err
	}
	return true, err
}

func retrievePostDetails(postReq GetPostRequest) ([]PostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var posts []PostDBModel
	var postResp []PostResponse
	filter := bson.M{"$and": []bson.M{{"username": postReq.UserName}, {"community_id": postReq.CommunityID}}} //bson.D{primitive.E{Key: "community_id", Value: postReq.CommunityID}, primitive.E{Key: "username", Value: postReq.UserName}}
	cursor, err := PostsCollection.Find(ctx, filter)
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

func retrieveFeedDetails(feedReq GetFeedRequest) ([]PostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var posts []PostDBModel
	var postResp []PostResponse
	feedFilter := bson.D{primitive.E{}}
	feedOptions := options.Find()
	if feedReq.Mode == "latest" {
		//fmt.Println("in feedReq.Mode == 'latest'")
		feedOptions.SetSort(bson.M{"created_at": -1})
	}
	if feedReq.Mode == "hot" {
		RankMostPosts()
		feedOptions.SetSort(bson.M{"ranking": -1})
	}
	if feedReq.PageNumber > 0 {
		//fmt.Println("in feedReq.PageNumber > 0")
		feedOptions.SetSkip(int64((feedReq.PageNumber - 1) * feedReq.NumberOfPosts))
	}
	if feedReq.PageNumber < 1 {
		//fmt.Println("in feedReq.PageNumber < 1")
		feedOptions.SetSkip(0)
	}
	if feedReq.NumberOfPosts > 0 {
		feedOptions.SetLimit(int64(feedReq.NumberOfPosts))
	}
	if feedReq.NumberOfPosts < 1 {
		//fmt.Println("in feedReq.NumberOfPosts < 1")
		feedOptions.SetLimit(10)
	}
	cursor, err := PostsCollection.Find(ctx, feedFilter, feedOptions)
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

func deletePost(postReq DeletePostRequest) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//filter := bson.D{primitive.E{Key: "username", Value: postReq.UserName}}
	filter := bson.M{"$and": []bson.M{{"username": postReq.UserName}, {"_id": postReq.ID}}}
	result, err := PostsCollection.DeleteOne(ctx, filter)
	return result, err
}

func editPostDetails(postReq EditPostRequest) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// updating the data in db
	delPostReq, err := ConvertEditPostReqToDeletePostReq(postReq)
	postExists, err := CheckPostExists(delPostReq)
	if !postExists {
		return result, err
	}
	filter := bson.M{"$and": []bson.M{{"username": postReq.UserName}, {"_id": postReq.ID}}}
	updateQuery := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "title", Value: postReq.Title},
				primitive.E{Key: "body", Value: postReq.Body},
			},
		},
	}
	result, err = PostsCollection.UpdateOne(ctx, filter, updateQuery)
	return result, err
}

// func RankPosts(feedReq GetFeedRequest) ([]PostResponse, error) {
// 	// formula
// 	// ranking = ( votes + comments / 3 ) / ( age_minutes + 120 )
// 	Completed, err := RankMostPosts()

// }

func RankMostPosts() (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var posts []PostDBModel
	//var postResp []PostResponse
	feedFilter := bson.D{primitive.E{}}
	feedOptions := options.Find()
	feedOptions.SetLimit(100)
	feedOptions.SetSort(bson.M{"upvotes": -1})
	cursor, err := PostsCollection.Find(ctx, feedFilter, feedOptions)
	if err = cursor.All(ctx, &posts); err != nil {
		return result, err
	}
	for _, post := range posts {
		filter := bson.M{"$and": []bson.M{{"username": post.UserName}, {"_id": post.ID}}}
		updateQuery := bson.D{
			{"$set", bson.D{{"ranking", UpdatePostRanking(post)}}},
		}
		//post := UpdatePostRanking(post, rank)
		//item, err := ConvertPostDBModelToPostResponse(post)
		result, err = PostsCollection.UpdateOne(ctx, filter, updateQuery)
		if err != nil {
			return result, err
		}
	}
	return result, err
}
func Routes(router *gin.Engine) {
	router.POST(POST_ROUTE_PREFIX, CreatePost())
	router.GET(POST_ROUTE_PREFIX, GetPosts())
	router.GET(HOME_ROUTE_PREFIX, GetFeed())
	router.DELETE(POST_ROUTE_PREFIX, DeletePost())
	router.PATCH(POST_ROUTE_PREFIX, EditPost())
}
