package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/Nikolat27/king_of_numbers/backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (handler *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 1 KB
	utils.ParseJSON(w, r.Body, 1<<10, &input)

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.WriteErrorJSON(w, 400, err.Error())
	}

	filter := bson.M{
		"username": input.Username,
	}

	if _, err := handler.Models.User.Get(filter); err == nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "username already taken")
		return
	} else if !errors.Is(err, mongo.ErrNoDocuments) {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, "database error: "+err.Error())
		return
	}

	if _, err := handler.Models.User.Create(input.Username, hashedPassword); err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, "could not create user: "+err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "user registered successfully")
}

func (handler *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 1 KB
	utils.ParseJSON(w, r.Body, 1<<10, &input)

	filter := bson.M{
		"username": input.Username,
	}

	userInstance, err := handler.Models.User.Get(filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			utils.WriteErrorJSON(w, http.StatusBadRequest, "password or username is incorrect")
			return
		}

		utils.WriteErrorJSON(w, http.StatusBadRequest, "could not get the user"+err.Error())
		return
	}

	isPasswordValid := utils.VerifyPassword(userInstance.Password, input.Password)
	if !isPasswordValid {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "password or username is incorrect")
		return
	}

	authToken, err := handler.Paseto.CreateToken(userInstance.Id, input.Username, 12*time.Hour)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, "failed to create token"+err.Error())
		return
	}

	response := map[string]string{
		"token": authToken,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}
