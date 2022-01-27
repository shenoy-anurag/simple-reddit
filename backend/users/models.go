package users

import (
	"time"
)

// Multiple models in order to reduce coupling, implement single responsibility, and to emulate private class variables in Go.
type CreateUserRequest struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	Name     string    `json:"name,omitempty"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Joined   time.Time `bson:"joined"`
}

type UserDBModel struct {
	Name     string    `json:"name,omitempty"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Joined   time.Time `bson:"joined"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Convertion functions to convert between different models.
func ConvertUserRequestToUserDBModel(usrReq CreateUserRequest) UserDBModel {
	return UserDBModel{
		Name:     usrReq.Name,
		Email:    usrReq.Email,
		Username: usrReq.Username,
		Password: usrReq.Password,
		Joined:   time.Now().UTC(),
	}
}
