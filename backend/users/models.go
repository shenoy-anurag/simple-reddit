package users

import (
	"time"
	"simple-reddit/profiles"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Multiple models in order to reduce coupling, and implement single responsibility.
type CreateUserRequest struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type UserResponse struct {
	FirstName string    `json:"firstname,omitempty"`
	LastName  string    `json:"lastname,omitempty"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Joined    time.Time `bson:"joined"`
}

type UserDBModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `json:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty"`
	Email     string             `json:"email"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Subcriptions []primitive.ObjectID `json:"subcriptions"`
	Joined    time.Time          `bson:"joined"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CheckUsernameRequest struct {
	Username string `json:"username" validate:"required"`
}

type GetSubsciptionsRequest struct {
	Username string `json:"username" validate:"required"`
}

type UpdateSubsciptionRequest struct {
	Username string `json:"username" validate:"required"`
	CommunityName        string `json:"communityname" validate:"required"`
}

type CommunityDBModel struct {
	ID              primitive.ObjectID `bson:"_id"`
	UserName        string             `json:"username" validate:"required"`
	Name            string             `bson:"name"`
	Description     string             `bson:"description"`
	SubscriberCount int                `bson:"subscriber_count"`
	CreatedAt       time.Time          `bson:"created_at"`
}

// Convertion functions to convert between different models.
func ConvertUserRequestToUserDBModel(usrReq CreateUserRequest) UserDBModel {
	return UserDBModel{
		ID:        primitive.NewObjectID(),
		FirstName: usrReq.FirstName,
		LastName:  usrReq.LastName,
		Email:     usrReq.Email,
		Username:  usrReq.Username,
		Password:  usrReq.Password,
		Subcriptions: []primitive.ObjectID{},// CreateSubcriptions(),
		Joined:    time.Now().UTC(),
	}
}

func ConvertUserDBModelToUserResponse(userDB UserDBModel) UserResponse {
	return UserResponse{
		FirstName: userDB.FirstName,
		LastName:  userDB.LastName,
		Email:     userDB.Email,
		Username:  userDB.Username,
		Joined:    userDB.Joined,
	}
}

func ConvertUserDBModelToProfileDBModel(userDB UserDBModel) profiles.ProfileDBModel {
	return profiles.ProfileDBModel{
		ID:        userDB.ID,
		FirstName: userDB.FirstName,
		LastName:  userDB.LastName,
		Email:     userDB.Email,
		UserName:  userDB.Username,
		Karma:     0,
		SavedPC: profiles.CreateSavedPC(userDB.Username),
		Birthday:  userDB.Joined,
	}
}
