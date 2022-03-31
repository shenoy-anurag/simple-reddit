package comments

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	DeletedAt time.Time          `bson:"deleted_at"`
}

type CreateCommentRequest struct {
	PostId   string `json:"post_id" validate:"required"`
	UserId   string `json:"user_id" validate:"required"`
	ParentId string `json:"parent_id" validate:"omitempty"`
	Body     string `json:"body" validate:"required"`
}

type DeleteCommentRequest struct {
	CommentId string `json:"comment_id" uri:"comment_id" validate:"required"`
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
		DeletedAt: time.Time{},
	}, nil
}

// Model CRUD functions
func createCommentInDB(comment CreateCommentRequest) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newComment, err := ConvertCommentRequestToCommentDBModel(comment) //, err := ConvertCommunityRequestToCommunityDBModel(community)
	if err != nil {
		return result, err
	}
	result, err = CommentsCollection.InsertOne(ctx, newComment)
	return result, err
}

func retrieveCommentById(comment_id_hex string) (CommentDBModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var comment CommentDBModel
	comment_id, err := primitive.ObjectIDFromHex(comment_id_hex)
	if err != nil {
		return comment, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: comment_id}}
	err = CommentsCollection.FindOne(ctx, filter).Decode(&comment)
	return comment, err
}

func deleteComment(commReq DeleteCommentRequest) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	comment, err := retrieveCommentById(commReq.CommentId)
	if err != nil {
		return result, err
	}
	// updating the comment in db
	filter := bson.M{"_id": comment.ID}
	updateQuery := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "body", Value: ""},
				primitive.E{Key: "is_deleted", Value: true},
				primitive.E{Key: "is_votable", Value: false},
				primitive.E{Key: "deleted_at", Value: time.Now().UTC()},
			},
		},
	}
	result, err = CommentsCollection.UpdateOne(
		ctx,
		filter,
		updateQuery,
	)
	return result, err
}
