package users

import (
	"time"

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
	Joined    time.Time          `bson:"joined"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CheckUsernameRequest struct {
	Username string `json:"username" validate:"required"`
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
