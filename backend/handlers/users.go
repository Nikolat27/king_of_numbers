package handlers

import (
	"fmt"
	"net/http"

	"github.com/Nikolat27/king_of_numbers/backend/utils"
)

func (handler *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 10 KB
	utils.ParseJSON(w, r.Body, 1<<10, &input)

	fmt.Println(input.Password)
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.WriteErrorJSON(w, 400, err.Error())
	}

	fmt.Printf("%s", hashedPassword)
}
