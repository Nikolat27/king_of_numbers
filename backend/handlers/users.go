package handlers

import (
	"net/http"

	"github.com/Nikolat27/king_of_numbers/backend/types"
	"github.com/Nikolat27/king_of_numbers/backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	UserExists = "userExists"
)

func (handler *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 10 KB
	utils.ParseJSON(w, r.Body, 1<<10, &input)

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	if err := checkUserExists(handler, input.Username); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	userId, err := handler.Models.User.Create(input.Username, hashedPassword)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	resp := map[string]string{
		"data":    "user registered successfully",
		"user_id": userId.Hex(),
	}
	utils.WriteJSON(w, http.StatusCreated, resp)
}

func checkUserExists(handler *Handler, username string) *types.ErrorResponse {
	filter := bson.M{
		"username": username,
	}

	projection := bson.M{
		"_id": 1,
	}

	if _, err := handler.Models.User.Get(filter, projection); err != nil {
		// if user does not exist, don`t return an error
		if err.Detail == mongo.ErrNoDocuments.Error() {
			return nil
		}

		// if error was not about existence (like internal database errors)
		return err
	}

	// if we did not have any error (which means user exists)
	return types.NewErrorResponse(UserExists, "user with this username exists")
}
