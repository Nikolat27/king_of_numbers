package utils

import (
	"errors"
	"net/http"

	"github.com/Nikolat27/king_of_numbers/backend/paseto"
)

func CheckAuth(r *http.Request, paseto *paseto.Maker) (*paseto.Payload, error) {
	cookie, err := r.Cookie("auth_cookie")
	if err != nil {
		return nil, errors.New("auth cookie is missing")
	}

	payload, err := paseto.VerifyToken(cookie.Value)
	if err != nil {
		return nil, errors.New("can't verify auth token")
	}

	return payload, nil
}
