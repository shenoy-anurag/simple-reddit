package comments

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Models for Comments
type CommentDBModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	PostId    primitive.ObjectID `bson:"post_id" validate:"required"`
	UserId    primitive.ObjectID `bson:"user_id" validate:"required"`
	ParentId  primitive.ObjectID `bson:"parent_id"`
	Body      string             `bson:"body" validate:"required"`
	Upvotes   int                `bson:"upvotes"`
	Downvotes int                `bson:"downvotes"`
	IsRoot    bool               `bson:"is_root"`
	IsVotable bool               `bson:"is_votable"`
	IsDeleted bool               `bson:"is_deleted"`
	CreatedAt time.Time          `bson:"created_at"`
}

type CreateCommentRequest struct {
	PostId   string `json:"post_id" validate:"required"`
	UserId   string `json:"user_id" validate:"required"`
	ParentId string `json:"parent_id" validate:"omitempty"`
	Body     string `json:"body" validate:"required"`
}

// Convertion functions to convert between different models.
func ConvertCommentRequestToCommentDBModel(commentReq CreateCommentRequest) (CommentDBModel, error) {
	commDBModel := CommentDBModel{}
	post_id, err := primitive.ObjectIDFromHex(commentReq.PostId)
	if err != nil {
		return commDBModel, err
	}
	user_id, err := primitive.ObjectIDFromHex(commentReq.UserId)
	if err != nil {
		return commDBModel, err
	}
	parent_id, err := primitive.ObjectIDFromHex(commentReq.ParentId)
	var is_root bool = false
	if err != nil {
		is_root = true
	}
	return CommentDBModel{
		ID:        primitive.NewObjectID(),
		PostId:    post_id,
		UserId:    user_id,
		ParentId:  parent_id,
		Body:      commentReq.Body,
		Upvotes:   0,
		Downvotes: 0,
		IsRoot:    is_root,
		IsVotable: true,
		IsDeleted: false,
		CreatedAt: time.Now().UTC(),
	}, nil
}
