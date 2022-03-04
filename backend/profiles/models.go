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
	Birthday  time.Time          `bson:"birthday"`
}

type GetProfileRequest struct {
	UserName    string             `json:"username" validate:"required"`
}

type ProfileResponse struct {
	FirstName string    `json:"firstname,omitempty"`
	LastName  string    `json:"lastname,omitempty"`
	Email     string    `json:"email"`
	UserName  string    `json:"username"`
	Birthday    time.Time `json:"joined"`
}

func ConvertProfileDBModelToProfileResponse(profileDB  ProfileDBModel) ProfileResponse {
	return ProfileResponse{
		FirstName: profileDB.FirstName,
		LastName:  profileDB.LastName,
		Email:     profileDB.Email,
		UserName:  profileDB.UserName,
		Birthday: profileDB.Birthday,
	}
}
