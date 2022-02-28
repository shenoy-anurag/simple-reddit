package users

import (
	"time"
)

// Multiple models in order to reduce coupling, implement single responsibility,
// and to emulate private class variables in Go.
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
	FirstName string    `json:"firstname,omitempty"`
	LastName  string    `json:"lastname,omitempty"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Joined    time.Time `bson:"joined"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CheckUsernameRequest struct {
	Username string `json:"username" validate:"required"`
}

// type ProfileDBModel struct {
// }

// Convertion functions to convert between different models.
func ConvertUserRequestToUserDBModel(usrReq CreateUserRequest) UserDBModel {
	return UserDBModel{
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
