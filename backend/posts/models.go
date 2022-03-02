package posts

import (
	"simple-reddit/users"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Multiple models in order to reduce coupling, implement single responsibility,
// and to emulate private class variables in Go.
type PostDBModel struct {
	CommunityID primitive.ObjectID `bson:"community_id"`
	UserID      primitive.ObjectID `bson:"user_id"`
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

// Convertion functions to convert between different models.
func ConvertPostRequestToPostDBModel(postReq CreatePostRequest) (PostDBModel, error) {
	userDB, err := users.GetUserDetails(postReq.UserName)
	return PostDBModel{
		CommunityID: postReq.CommunityID,
		UserID:      userDB.ID,
		Title:       postReq.Title,
		Body:        postReq.Body,
		Upvotes:     0,
		Downvotes:   0,
		CreatedAt:   time.Now().UTC(),
	}, err
}
