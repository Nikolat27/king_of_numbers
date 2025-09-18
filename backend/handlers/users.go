package handlers

import (
	"net/http"
	"time"

	"github.com/Nikolat27/king_of_numbers/backend/database/models"
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

	isExist, _, err := checkUserExists(handler, input.Username)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	if isExist {
		utils.WriteErrorJSON(w, http.StatusBadRequest,
			types.NewErrorResponse("checkUser", "user with this username exists"))

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

func (handler *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 10 KB
	utils.ParseJSON(w, r.Body, 1<<10, &input)

	isExist, userInstance, err := checkUserExists(handler, input.Username)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	if !isExist {
		utils.WriteErrorJSON(w, http.StatusBadRequest,
			types.NewErrorResponse("verifyCredentials", "username or password is incorrect"))

		return
	}

	if !utils.VerifyPassword(userInstance.Password, input.Password) {
		utils.WriteErrorJSON(w, http.StatusBadRequest,
			types.NewErrorResponse("verifyCredentials", "username or password is incorrect"))

		return
	}

	authToken, err := handler.Paseto.CreateToken(userInstance.Id, userInstance.Username, 12*time.Hour)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	resp := map[string]string{
		"data":       "user logged in successfully",
		"auth_token": authToken,
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}

func checkUserExists(handler *Handler, username string) (bool, *models.User, *types.ErrorResponse) {
	filter := bson.M{
		"username": username,
	}

	projection := bson.M{
		"_id":      1,
		"username": 1,
		"password": 1,
	}

	userInstance, err := handler.Models.User.Get(filter, projection)
	if err != nil {
		// if user does not exist, don`t return an error
		if err.Detail == mongo.ErrNoDocuments.Error() {
			return false, &models.User{}, nil
		}

		// if error was not about existence (like internal database errors)
		return false, &models.User{}, err
	}

	// if we did not have any error (which means user exists)
	return true, userInstance, nil
}
