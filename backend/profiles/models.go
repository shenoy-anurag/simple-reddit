package profiles

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileDBModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `json:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty"`
	Email     string             `json:"email"`
	UserName  string             `json:"username"`
	Karma     int                `json:"karma"`
	SavedPC   SavedDBModel       `json:"savedpc"`
	Birthday  time.Time          `bson:"birthday"`
}

type GetProfileRequest struct {
	UserName string `json:"username" validate:"required"`
}

type ProfileResponse struct {
	FirstName string    `json:"firstname,omitempty"`
	LastName  string    `json:"lastname,omitempty"`
	Email     string    `json:"email"`
	UserName  string    `json:"username"`
	Karma     int       `json:"karma"`
	Birthday  time.Time `json:"joined"`
}

type EditProfileRequest struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required"`
	UserName  string `json:"username" validate:"required"`
}

type DeleteProfileRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserDBModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `json:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty"`
	Email     string             `json:"email"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Joined    time.Time          `bson:"joined"`
}

type SavedDBModel struct {
	ID            primitive.ObjectID   `bson:"_id" validate:"required"`
	Username      string               `json:"username" validate:"required"`
	SavedPosts    []primitive.ObjectID `bson:"savedposts" validate:"required"`
	SavedComments []primitive.ObjectID `bson:"savedcomments" validate:"required"`
}

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

type PostResponse struct {
	ID        primitive.ObjectID `json:"_id"`
	Title     string             `json:"title"`
	Body      string             `json:"body"`
	Upvotes   int                `json:"upvotes"`
	Downvotes int                `json:"downvotes"`
	CreatedAt time.Time          `json:"created_at"`
}

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

type UpdateSavedPostRequest struct {
	PostID   primitive.ObjectID `json:"post_id" validate:"required"`
	Username string             `json:"username" validate:"required"`
}

type UpdateSavedCommentRequest struct {
	CommentID primitive.ObjectID `json:"comment_id" validate:"required"`
	Username  string             `json:"username" validate:"required"`
}

type GetSavedItemRequest struct {
	Username string `json:"username" validate:"required"`
}

// type GetSavedCommentRequest struct {
// 	Username string `json:"username" validate:"required"`
// }

// Convertion functions to convertbetween different models.

func ConvertEditProfileRequestToPrfileDBModel(editProfile EditProfileRequest) ProfileDBModel {
	return ProfileDBModel{
		FirstName: editProfile.FirstName,
		LastName:  editProfile.LastName,
		Email:     editProfile.Email,
		UserName:  editProfile.UserName,
	}
}

func ConvertProfileDBModelToProfileResponse(profileDB ProfileDBModel) ProfileResponse {
	return ProfileResponse{
		FirstName: profileDB.FirstName,
		LastName:  profileDB.LastName,
		Email:     profileDB.Email,
		UserName:  profileDB.UserName,
		Karma:     profileDB.Karma,
		Birthday:  profileDB.Birthday,
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

// func ConvertCVRToCVHDBModel(cVoteReq CommentVoteRequest) (CommentVoteHistoryDBModel, error) {
// 	cVHDbModel := CommentVoteHistoryDBModel{}
// 	newId := primitive.NewObjectID()
//
// 	comment_id, err := primitive.ObjectIDFromHex(cVoteReq.CommentId)
// 	if err != nil {
// 		return cVHDbModel, err
// 	}
// 	var is_upvote bool = false
// 	var is_downvote bool = false
// 	if cVoteReq.Vote == UPVOTE {
// 		is_upvote = true
// 	} else if cVoteReq.Vote == DOWNVOTE {
// 		is_downvote = true
// 	} else {
// 		return cVHDbModel, err
// 	}
// 	currTime := time.Now().UTC()
// 	return CommentVoteHistoryDBModel{
// 		ID:            newId,
// 		UserName:      cVoteReq.UserName,
// 		CommentId:     comment_id,
// 		IsUpvoted:     is_upvote,
// 		IsDownvoted:   is_downvote,
// 		CreatedAt:     currTime,
// 		LastUpdatedAt: currTime,
// 	}, nil
// }

func ConvertCommentDBModelToCommentResponse(comment CommentDBModel) CommentResponse {
	return CommentResponse(comment)
}

func CreateSavedDBModel(UserName string) SavedDBModel {
	return SavedDBModel{
		ID:            primitive.NewObjectID(),
		Username:      UserName,
		SavedPosts:    []primitive.ObjectID{},
		SavedComments: []primitive.ObjectID{},
	}
}
