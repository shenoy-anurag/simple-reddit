package communities

import (
	"time"
	// "github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Models for Communities
type CommunityDBModel struct {
	ID              primitive.ObjectID `bson:"_id"`
	UserName        string             `json:"username" validate:"required"`
	Name            string             `bson:"name"`
	Description     string             `bson:"description"`
	SubscriberCount int                `bson:"subscriber_count"`
	CreatedAt       time.Time          `bson:"created_at"`
}

type CreateCommunityRequest struct {
	UserName    string `json:"username" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type GetCommunityRequest struct {
	Name   string `json:"name" validate:"required"`
	IsUser bool   `json:"isuser" validate:"exists"`
}

type GetAllCommunitiesRequest struct {
	PageNumber   uint32 `json:"p" form:"p"`
	ItemsPerPage uint32 `json:"per_page" form:"per_page"`
}

// GetAllCommunitiesRequest struct constructor function
func (gcr *GetAllCommunitiesRequest) fill_defaults() {
	// setting default values
	// if no values present
	if gcr.PageNumber == 0 {
		gcr.PageNumber = 1
	}
	if gcr.ItemsPerPage == 0 {
		gcr.ItemsPerPage = 25
	}
}

type GetPostsRequest struct {
	Name string `json:"name" validate:"required"`
}

type EditCommunityRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type DeleteCommunityRequest struct {
	Name     string `json:"name" uri:"name" validate:"required"`
	UserName string `json:"username" uri:"username" validate:"required"`
}

type CommunityResponse struct {
	ID              primitive.ObjectID `json:"_id"`
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	SubscriberCount int                `json:"subscriber_count"`
	CreatedAt       time.Time          `json:"created_at"`
}

type GetFeedRequest struct {
	PageNumber    int    `json:"pagenumber"`
	NumberOfPosts int    `json:"numberofposts"`
	Mode          string `json:"mode"`
}

// use a single instance of Validate, it caches struct info
// var valdate *validator.Validate

// Validate validate input data
// func ValidateGetCommunityRequest(gcr *GetCommunityRequest) (Validate(), error) {
// 	valdate := validator.New()
// 	return valdate.Struct(gcr)
// }

// Convertion functions to convert between different models.
func ConvertCommunityRequestToCommunityDBModel(communityReq CreateCommunityRequest) CommunityDBModel { //, error) {
	//user_id, err := primitive.ObjectIDFromHex(communityReq.UserID)
	return CommunityDBModel{
		ID:              primitive.NewObjectID(),
		UserName:        communityReq.UserName,
		Name:            communityReq.Name,
		Description:     communityReq.Description,
		SubscriberCount: 0,
		CreatedAt:       time.Now().UTC(),
	} //, err
}

func ConvertCommunityDBModelToCommunityResponse(communityDB CommunityDBModel) CommunityResponse {
	return CommunityResponse{
		ID:              communityDB.ID,
		Name:            communityDB.Name,
		Description:     communityDB.Description,
		SubscriberCount: communityDB.SubscriberCount,
		CreatedAt:       communityDB.CreatedAt,
	}
}

func ConvertEditCommunityReqToGetCommunityReq(communityReq EditCommunityRequest) GetCommunityRequest {
	return GetCommunityRequest{
		Name: communityReq.Name,
	}
}

type PostDBModel struct {
	ID          primitive.ObjectID `bson:"_id"`
	CommunityID primitive.ObjectID `bson:"community_id"`
	UserName    string             `bson:"username"`
	Title       string             `bson:"title"`
	Body        string             `json:"body"`
	Upvotes     int                `bson:"upvotes"`
	Downvotes   int                `bson:"downvotes"`
	CreatedAt   time.Time          `bson:"created_at"`
}

type PostResponse struct {
	ID        primitive.ObjectID `json:"_id"`
	Title     string             `json:"title"`
	Body      string             `json:"body"`
	Upvotes   int                `json:"upvotes"`
	Downvotes int                `json:"downvotes"`
	CreatedAt time.Time          `json:"created_at"`
}

func ConvertPostDBModelToPostResponse(postDB PostDBModel) (PostResponse, error) {
	var err error
	return PostResponse{
		ID:        postDB.ID,
		Title:     postDB.Title,
		Body:      postDB.Body,
		Upvotes:   postDB.Upvotes,
		Downvotes: postDB.Downvotes,
		CreatedAt: postDB.CreatedAt,
	}, err
}

type CommunitySubscriberDBModel struct {
	ID          primitive.ObjectID `bson:"_id"`
	UserID      primitive.ObjectID `bson:"user_id"`
	CommunityID primitive.ObjectID `bson:"community_id"`
	JoinedAt    time.Time          `bson:"joined_at"`
}
