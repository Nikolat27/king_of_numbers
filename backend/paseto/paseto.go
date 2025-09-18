package paseto

import (
	"crypto/sha256"
	"os"
	"time"

	"github.com/Nikolat27/king_of_numbers/backend/types"
	"github.com/o1egl/paseto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Maker struct {
	paseto       *paseto.V2
	symmetricKey [32]byte
}

func New() (*Maker, *types.ErrorResponse) {
	key, errResponse := getSymmetricKey()
	if errResponse != nil {
		return nil, errResponse
	}

	maker := &Maker{
		paseto:       paseto.NewV2(),
		symmetricKey: key,
	}

	return maker, nil
}

func getSymmetricKey() ([32]byte, *types.ErrorResponse) {
	symmetricKey := os.Getenv("PASETO_SYMMETRIC_KEY")
	if symmetricKey == "" {
		return [32]byte{}, types.NewErrorResponse(".env var does not exist",
			"PASETO_SYMMETRIC_KEY env variable does not exist")
	}

	hash := sha256.Sum256([]byte(symmetricKey))

	return hash, &types.ErrorResponse{}
}

func (maker *Maker) CreateToken(userId primitive.ObjectID, username string,
	duration time.Duration) (string, *types.ErrorResponse) {

	payload, err := NewPayload(userId, username, duration)
	if err != nil {
		return "", types.NewErrorResponse("new auth token", err.Error())
	}
	token, err := maker.paseto.Encrypt(maker.symmetricKey[:], payload, nil)
	if err != nil {
		types.NewErrorResponse("generate auth token", err.Error())
	}

	return token, nil
}

func (maker *Maker) VerifyToken(token string) (*Payload, *types.ErrorResponse) {
	var payload Payload
	if err := maker.paseto.Decrypt(token, maker.symmetricKey[:], &payload, nil); err != nil {
		return nil, types.NewErrorResponse("verify auth token", err.Error())
	}

	if err := payload.Valid(); err != nil {
		return nil, types.NewErrorResponse("verify auth token", err.Error())
	}

	return &payload, nil
}
