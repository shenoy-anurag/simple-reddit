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

// Convertion functions to convert between different models.

func ConvertEditProfileRequestToProfileDBModel(editProfile EditProfileRequest) ProfileDBModel {
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
