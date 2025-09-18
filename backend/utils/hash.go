package utils

import (
	"github.com/Nikolat27/king_of_numbers/backend/types"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, *types.ErrorResponse) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", types.NewErrorResponse("hashPassword", err)
	}

	return string(hash), nil
}

func VerifyPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
