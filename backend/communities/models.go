package communities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Models for Communities
type CommunityDBModel struct {
	ID              primitive.ObjectID `bson:"_id"`
	UserID          primitive.ObjectID `bson:"user_id"`
	Name            string             `bson:"name"`
	Description     string             `bson:"description"`
	SubscriberCount int                `bson:"subscriber_count"`
	CreatedAt       time.Time          `bson:"created_at"`
}

type CreateCommunityRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// Convertion functions to convert between different models.
func ConvertCommunityRequestToCommunityDBModel(communityReq CreateCommunityRequest) (CommunityDBModel, error) {
	user_id, err := primitive.ObjectIDFromHex(communityReq.UserID)
	return CommunityDBModel{
		ID:              primitive.NewObjectID(),
		UserID:          user_id,
		Name:            communityReq.Name,
		Description:     communityReq.Description,
		SubscriberCount: 0,
		CreatedAt:       time.Now().UTC(),
	}, err
}

type CommunitySubscriberDBModel struct {
	ID          primitive.ObjectID `bson:"_id"`
	UserID      primitive.ObjectID `bson:"user_id"`
	CommunityID primitive.ObjectID `bson:"community_id"`
	JoinedAt    time.Time          `bson:"joined_at"`
}
