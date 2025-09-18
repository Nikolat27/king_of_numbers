package utils

import (
	"net/http"

	"github.com/Nikolat27/king_of_numbers/backend/paseto"
	"github.com/Nikolat27/king_of_numbers/backend/types"
)

func CheckAuth(r *http.Request, paseto *paseto.Maker) (*paseto.Payload, *types.ErrorResponse) {
	cookie, err := r.Cookie("auth_cookie")
	if err != nil {
		return nil, types.NewErrorResponse("authCookie", err)
	}

	payload, errResp := paseto.VerifyToken(cookie.Value)
	if errResp != nil {
		return nil, errResp
	}

	return payload, nil
}
