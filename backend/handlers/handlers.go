package handlers

import (
	"github.com/Nikolat27/king_of_numbers/backend/database/models"
	"github.com/Nikolat27/king_of_numbers/backend/paseto"
)

type Handler struct {
	Models *models.Models
	Paseto *paseto.Maker
}

// New -> Construction
func New(models *models.Models, paseto *paseto.Maker) *Handler {
	return &Handler{
		Models: models,
		Paseto: paseto,
	}
}
