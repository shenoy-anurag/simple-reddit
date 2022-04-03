package comments

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UPVOTE string = "upvote"
const DOWNVOTE string = "downvote"

// Models for Comments
type CommentDBModel struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserName   string             `bson:"username"`
	PostId     primitive.ObjectID `bson:"post_id" validate:"required"`
	ParentId   primitive.ObjectID `bson:"parent_id"`
	Body       string             `bson:"body" validate:"required"`
	Upvotes    int                `bson:"upvotes"`
	Downvotes  int                `bson:"downvotes"`
	TotalVotes int                `bson:"total_votes"`
	IsRoot     bool               `bson:"is_root"`
	IsVotable  bool               `bson:"is_votable"`
	IsDeleted  bool               `bson:"is_deleted"`
	CreatedAt  time.Time          `bson:"created_at"`
	DeletedAt  time.Time          `bson:"deleted_at"`
}

type CreateCommentRequest struct {
	UserName string `json:"username" validate:"required"`
	PostId   string `json:"post_id" validate:"required"`
	ParentId string `json:"parent_id" validate:"omitempty"`
	Body     string `json:"body" validate:"required"`
}

type DeleteCommentRequest struct {
	CommentId string `json:"comment_id" uri:"comment_id" form:"comment_id" validate:"required"`
}

type GetCommentRequest struct {
	PostId string `json:"post_id" uri:"post_id" form:"post_id" validate:"required"`
}

type CommentResponse struct {
	ID         primitive.ObjectID `json:"_id"`
	UserName   string             `json:"username"`
	PostId     primitive.ObjectID `json:"post_id"`
	ParentId   primitive.ObjectID `json:"parent_id"`
	Body       string             `json:"body"`
	Upvotes    int                `json:"upvotes"`
	Downvotes  int                `json:"downvotes"`
	TotalVotes int                `json:"total_votes"`
	IsRoot     bool               `json:"is_root"`
	IsVotable  bool               `json:"is_votable"`
	IsDeleted  bool               `json:"is_deleted"`
	CreatedAt  time.Time          `json:"created_at"`
	DeletedAt  time.Time          `json:"deleted_at"`
}

// Models for Comment Voting History (to prevent a single user from giving multiple upvotes/downvotes on the same comment)
type CommentVoteHistoryDBModel struct {
	ID            primitive.ObjectID `bson:"_id"`
	UserName      string             `bson:"username"`
	CommentId     primitive.ObjectID `bson:"comment_id"`
	IsUpvoted     bool               `bson:"is_upvoted"`
	IsDownvoted   bool               `bson:"is_downvoted"`
	CreatedAt     time.Time          `bson:"created_at"`
	LastUpdatedAt time.Time          `bson:"last_updated_at"`
}

type CommentVoteRequest struct {
	UserName  string `json:"username" validate:"required"`
	CommentId string `json:"comment_id" validate:"required"`
	Vote      string `json:"vote" validate:"required"`
}

// Convertion functions to convert between different models.
func ConvertCommentRequestToCommentDBModel(commentReq CreateCommentRequest) (CommentDBModel, error) {
	commDBModel := CommentDBModel{}
	post_id, err := primitive.ObjectIDFromHex(commentReq.PostId)
	if err != nil {
		return commDBModel, err
	}
	parent_id, err := primitive.ObjectIDFromHex(commentReq.ParentId)
	var is_root bool = false
	if err != nil {
		is_root = true
	}
	return CommentDBModel{
		ID:         primitive.NewObjectID(),
		UserName:   commentReq.UserName,
		PostId:     post_id,
		ParentId:   parent_id,
		Body:       commentReq.Body,
		Upvotes:    0,
		Downvotes:  0,
		TotalVotes: 0,
		IsRoot:     is_root,
		IsVotable:  true,
		IsDeleted:  false,
		CreatedAt:  time.Now().UTC(),
		DeletedAt:  time.Time{},
	}, nil
}

func ConvertCVRToCVHDBModel(cVoteReq CommentVoteRequest) (CommentVoteHistoryDBModel, error) {
	cVHDbModel := CommentVoteHistoryDBModel{}
	newId := primitive.NewObjectID()

	comment_id, err := primitive.ObjectIDFromHex(cVoteReq.CommentId)
	if err != nil {
		return cVHDbModel, err
	}
	var is_upvote bool = false
	var is_downvote bool = false
	if cVoteReq.Vote == UPVOTE {
		is_upvote = true
	} else if cVoteReq.Vote == DOWNVOTE {
		is_downvote = true
	} else {
		return cVHDbModel, err
	}
	currTime := time.Now().UTC()
	return CommentVoteHistoryDBModel{
		ID:            newId,
		UserName:      cVoteReq.UserName,
		CommentId:     comment_id,
		IsUpvoted:     is_upvote,
		IsDownvoted:   is_downvote,
		CreatedAt:     currTime,
		LastUpdatedAt: currTime,
	}, nil
}

func ConvertCommentDBModelToCommentResponse(comment CommentDBModel) CommentResponse {
	return CommentResponse(comment)
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

func retrieveAllCommentsOfPost(post_id_hex string) ([]CommentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var comments []CommentDBModel
	var commentResponses []CommentResponse
	post_id, err := primitive.ObjectIDFromHex(post_id_hex)
	if err != nil {
		return commentResponses, err
	}
	filter := bson.D{primitive.E{Key: "post_id", Value: post_id}}
	cursor, err := CommentsCollection.Find(ctx, filter)
	if err != nil {
		return commentResponses, err
	}
	if err = cursor.All(ctx, &comments); err != nil {
		return commentResponses, err
	}
	for _, comment := range comments {
		item := ConvertCommentDBModelToCommentResponse(comment)
		commentResponses = append(commentResponses, item)
	}
	return commentResponses, err
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

func updateVoteComment(comment_id_hex string, is_upvote bool, is_remove_vote bool) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// updating the comment in db
	comment_id, err := primitive.ObjectIDFromHex(comment_id_hex)
	if err != nil {
		return result, err
	}
	filter := bson.M{"_id": comment_id}
	updateQuery := bson.D{}
	if is_remove_vote {
		if is_upvote {
			updateQuery = bson.D{
				primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "total_votes", Value: 1}, primitive.E{Key: "downvotes", Value: -1}}},
			}
		} else {
			updateQuery = bson.D{
				primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "total_votes", Value: -1}, primitive.E{Key: "upvotes", Value: -1}}},
			}
		}
	} else if is_upvote {
		updateQuery = bson.D{
			primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "upvotes", Value: 1}, primitive.E{Key: "total_votes", Value: 1}}},
		}
	} else {
		updateQuery = bson.D{
			primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "downvotes", Value: 1}, primitive.E{Key: "total_votes", Value: -1}}},
		}
	}
	result, err = CommentsCollection.UpdateOne(
		ctx,
		filter,
		updateQuery,
	)
	return result, err
}

func updateVote(cVoteReq CommentVoteRequest) (result string, err error) {
	commentHist, err := retrieveCommentVoteHistForUser(cVoteReq.CommentId, cVoteReq.UserName)
	if err != nil && err != mongo.ErrNoDocuments {
		return "voted", err
	}

	var is_upvote bool = false
	var is_remove_vote bool = false
	if cVoteReq.Vote == UPVOTE {
		if commentHist.CommentId == primitive.NilObjectID {
			is_upvote = true
		} else if commentHist.IsUpvoted {
			return "voted", nil
		} else if commentHist.IsDownvoted {
			is_remove_vote = true
			is_upvote = true
		}
	} else if cVoteReq.Vote == DOWNVOTE {
		if commentHist.CommentId == primitive.NilObjectID {
			is_upvote = false
		} else if commentHist.IsDownvoted {
			return "voted", nil
		} else if commentHist.IsUpvoted {
			is_remove_vote = true
			is_upvote = false
		}
	}

	_, err = updateVoteComment(cVoteReq.CommentId, is_upvote, is_remove_vote)
	if err != nil {
		return "error", err
	}

	if is_remove_vote {
		_, err := deleteCommentVoteHistForUser(cVoteReq.CommentId, cVoteReq.UserName)
		return "deleted", err
	} else {
		cVoteHist, err := ConvertCVRToCVHDBModel(cVoteReq)
		if err != nil {
			return "error", err
		}
		_, err = createCommentVoteHistInDB(cVoteHist)
		if err != nil {
			return "error", err
		}
	}
	return "voted", nil
}

func createCommentVoteHistInDB(cVoteHist CommentVoteHistoryDBModel) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err = CommentsVotingHistoryCollection.InsertOne(ctx, cVoteHist)
	return result, err
}

func retrieveCommentVoteHistForUser(comment_id_hex string, username string) (CommentVoteHistoryDBModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var commentHist CommentVoteHistoryDBModel
	comment_id, err := primitive.ObjectIDFromHex(comment_id_hex)
	if err != nil {
		return commentHist, err
	}
	filter := bson.D{
		primitive.E{Key: "username", Value: username},
		primitive.E{Key: "comment_id", Value: comment_id},
	}
	err = CommentsVotingHistoryCollection.FindOne(ctx, filter).Decode(&commentHist)
	return commentHist, err
}

func deleteCommentVoteHistForUser(comment_id_hex string, username string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	comment_id, err := primitive.ObjectIDFromHex(comment_id_hex)
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		primitive.E{Key: "username", Value: username},
		primitive.E{Key: "comment_id", Value: comment_id},
	}
	result, err := CommentsVotingHistoryCollection.DeleteOne(ctx, filter)
	return result, err
}
