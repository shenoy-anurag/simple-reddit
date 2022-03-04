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
	Username  string             `json:"username"`
	Karma     int                `json:"karma"`
	Birthday  time.Time          `bson:"birthday"`
}
