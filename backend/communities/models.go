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

type GetCommunityRequest struct {
	Name string `json:"name" uri:"name" validate:"required"`
}

type EditCommunityRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type DeleteCommunityRequest struct {
	Name     string `json:"name" uri:"name" validate:"required"`
	UserName string `json:"username" uri:"username" validate:"required"`
}

type CommunityResponse struct {
	ID              primitive.ObjectID `json:"_id"`
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	SubscriberCount int                `json:"subscriber_count"`
	CreatedAt       time.Time          `json:"created_at"`
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

func ConvertCommunityDBModelToCommunityResponse(communityDB CommunityDBModel) CommunityResponse {
	return CommunityResponse{
		ID:              communityDB.ID,
		Name:            communityDB.Name,
		Description:     communityDB.Description,
		SubscriberCount: communityDB.SubscriberCount,
		CreatedAt:       communityDB.CreatedAt,
	}
}

func ConvertEditCommunityReqToGetCommunityReq(communityReq EditCommunityRequest) GetCommunityRequest {
	return GetCommunityRequest{
		Name: communityReq.Name,
	}
}

type CommunitySubscriberDBModel struct {
	ID          primitive.ObjectID `bson:"_id"`
	UserID      primitive.ObjectID `bson:"user_id"`
	CommunityID primitive.ObjectID `bson:"community_id"`
	JoinedAt    time.Time          `bson:"joined_at"`
}
