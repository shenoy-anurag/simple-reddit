package profiles

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProfileDBModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `json:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty"`
	Email     string             `json:"email"`
	UserName  string             `json:"username"`
	Karma     int                `json:"karma"`
	SavedPC   SavedDBModel	 `json:"savedpc"`
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
	ID         primitive.ObjectID `bson:"_id" validate:"required"`
	Username string `json:"username" validate:"required"`
	SavedPosts []primitive.ObjectID `bson:"savedposts" validate:"required"`
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

type UpdateSavedPostRequest struct {
	PostID primitive.ObjectID `json:"post_id" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type UpdateSavedCommentRequest struct {
	CommentID primitive.ObjectID `json:"comment_id" validate:"required"`
	Username string `json:"username" validate:"required"`
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

func CreateSavedDBModel(UserName string) SavedDBModel{
	return SavedDBModel{
		ID:        primitive.NewObjectID(),
		Username: UserName,
		SavedPosts: []primitive.ObjectID{},
		SavedComments: []primitive.ObjectID{},
	}
}
