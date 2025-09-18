package paseto

import (
	"time"

	"github.com/Nikolat27/king_of_numbers/backend/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payload struct {
	ID        uuid.UUID          `json:"id"`
	UserId    primitive.ObjectID `json:"user_id"`
	Username  string             `json:"username"`
	CreatedAt time.Time          `json:"created_at"`
	ExpiryAt  time.Time          `json:"expiry_at"`
}

func NewPayload(userId primitive.ObjectID, username string, duration time.Duration) (*Payload, *types.ErrorResponse) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, types.NewErrorResponse("random uuid", err)
	}

	payload := &Payload{
		ID:        tokenId,
		UserId:    userId,
		Username:  username,
		CreatedAt: time.Now(),
		ExpiryAt:  time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() *types.ErrorResponse {
	if time.Now().After(payload.ExpiryAt) {
		return types.NewErrorResponse("token validating", "token is expired")
	}

	return nil
}
