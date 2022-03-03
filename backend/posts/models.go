package posts

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Multiple models in order to reduce coupling, and implement single responsibility.
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

type CreatePostRequest struct {
	UserName    string             `json:"username" validate:"required"`
	CommunityID primitive.ObjectID `json:"community_id" validate:"required"`
	Title       string             `json:"title"`
	Body        string             `json:"body"`
}

type GetPostRequest struct {
	UserName    string             `json:"username" validate:"required"`
	CommunityID primitive.ObjectID `json:"community_id" validate:"required"`
}

type PostResponse struct {
	ID        primitive.ObjectID `json:"_id"`
	Title     string             `json:"title"`
	Body      string             `json:"body"`
	Upvotes   int                `json:"upvotes"`
	Downvotes int                `json:"downvotes"`
	CreatedAt time.Time          `json:"created_at"`
}

// Convertion functions to convert between different models.
func ConvertPostRequestToPostDBModel(postReq CreatePostRequest) PostDBModel {
	return PostDBModel{
		ID:          primitive.NewObjectID(),
		CommunityID: postReq.CommunityID,
		UserName:    postReq.UserName,
		Title:       postReq.Title,
		Body:        postReq.Body,
		Upvotes:     1,
		Downvotes:   0,
		CreatedAt:   time.Now().UTC(),
	}
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
