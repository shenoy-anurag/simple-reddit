package posts

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Multiple models in order to reduce coupling, and implement single responsibility.
type PostDBModel struct {
	ID          primitive.ObjectID `bson:"_id"`
	CommunityID primitive.ObjectID `bson:"community_id"`
	UserName    string             `bson:"username"`
	Title       string             `bson:"title"`
	Body        string             `bson:"body"`
	Upvotes     int                `bson:"upvotes"`
	Downvotes   int                `bson:"downvotes"`
	Ranking     int                `bson:"ranking"`
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

type DeletePostRequest struct {
	ID       primitive.ObjectID `json:"id" validate:"required"`
	UserName string             `json:"username" validate:"required"`
}

type EditPostRequest struct {
	ID       primitive.ObjectID `json:"id" validate:"required"`
	UserName string             `json:"username" validate:"required"`
	Title    string             `json:"title" validate:"required"`
	Body     string             `json:"body" validate:"required"`
}

type VoteRequest struct {
	ID       primitive.ObjectID `json:"id" validate:"required"`
	UserName string             `json:"username" validate:"required"`
	Vote     int                `json:"vote" validate:"required"`
}

// Models for Post Voting History (to prevent a single user from giving multiple upvotes/downvotes on the same comment)
type PostVoteHistoryDBModel struct {
	ID            primitive.ObjectID `bson:"_id"`
	UserName      string             `bson:"username"`
	PostId        primitive.ObjectID `bson:"post_id"`
	IsUpvoted     bool               `bson:"is_upvoted"`
	IsDownvoted   bool               `bson:"is_downvoted"`
	CreatedAt     time.Time          `bson:"created_at"`
	LastUpdatedAt time.Time          `bson:"last_updated_at"`
}

type GetFeedRequest struct {
	PageNumber    int    `json:"pagenumber"`
	NumberOfPosts int    `json:"numberofposts"`
	Mode          string `json:"mode"`
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
		Ranking:     0,
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

func ConvertEditPostReqToDeletePostReq(postReq EditPostRequest) (DeletePostRequest, error) {
	var err error
	return DeletePostRequest{
		ID:       postReq.ID,
		UserName: postReq.UserName,
	}, err
}

func ConvertVotePostReqToDeletePostReq(votereq VoteRequest) (DeletePostRequest, error) {
	var err error
	return DeletePostRequest{
		ID:       votereq.ID,
		UserName: votereq.UserName,
	}, err
}

func ConvertPVRToPVHDBModel(pVoteReq VoteRequest) (PostVoteHistoryDBModel, error) {
	pVHDbModel := PostVoteHistoryDBModel{}
	newId := primitive.NewObjectID()

	var is_upvote bool = false
	var is_downvote bool = false
	if pVoteReq.Vote > 0 {
		is_upvote = true
	} else if pVoteReq.Vote < 0 {
		is_downvote = true
	} else {
		return pVHDbModel, nil
	}
	currTime := time.Now().UTC()
	return PostVoteHistoryDBModel{
		ID:            newId,
		UserName:      pVoteReq.UserName,
		PostId:        pVoteReq.ID,
		IsUpvoted:     is_upvote,
		IsDownvoted:   is_downvote,
		CreatedAt:     currTime,
		LastUpdatedAt: currTime,
	}, nil
}

// CRUD Methods on models
func UpdatePostRanking(postDB PostDBModel) int {
	// ranking = ( votes + comments / 3 ) / ( age_minutes + 120 )
	rank := ((postDB.Upvotes - postDB.Downvotes) * 4) / (int(postDB.CreatedAt.Sub(time.Now().UTC())) * 100)
	return rank
}

func createPostVoteHistInDB(pVoteHist PostVoteHistoryDBModel) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err = PostsVotingHistoryCollection.InsertOne(ctx, pVoteHist)
	return result, err
}

func retrievePostVoteHistForUser(post_id primitive.ObjectID, username string) (PostVoteHistoryDBModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var postHist PostVoteHistoryDBModel
	filter := bson.D{
		primitive.E{Key: "username", Value: username},
		primitive.E{Key: "post_id", Value: post_id},
	}
	err := PostsVotingHistoryCollection.FindOne(ctx, filter).Decode(&postHist)
	return postHist, err
}

func updateVotePost(post_id primitive.ObjectID, is_upvote bool, is_remove_vote bool) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// updating the comment in db
	filter := bson.M{"_id": post_id}
	updateQuery := bson.D{}
	if is_remove_vote {
		if is_upvote {
			updateQuery = bson.D{
				primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "downvotes", Value: -1}}},
			}
		} else {
			updateQuery = bson.D{
				primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "upvotes", Value: -1}}},
			}
		}
	} else if is_upvote {
		updateQuery = bson.D{
			primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "upvotes", Value: 1}}},
		}
	} else {
		updateQuery = bson.D{
			primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "downvotes", Value: 1}}},
		}
	}
	result, err = PostsCollection.UpdateOne(
		ctx,
		filter,
		updateQuery,
	)
	return result, err
}

func deletePostVoteHistForUser(post_id primitive.ObjectID, username string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{
		primitive.E{Key: "username", Value: username},
		primitive.E{Key: "post_id", Value: post_id},
	}
	result, err := PostsVotingHistoryCollection.DeleteOne(ctx, filter)
	return result, err
}
