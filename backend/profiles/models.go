package profiles

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Multiple models in order to reduce coupling, and implement single responsibility.
type PostDBModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `json:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty"`
	Email     string             `json:"email"`
	Username  string             `json:"username"`
	Joined    time.Time          `bson:"joined"`
}

type ProfileRequest struct {
	Username string `json:"username"`
}
