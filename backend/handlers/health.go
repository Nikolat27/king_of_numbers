package handlers

import (
	"net/http"

	"github.com/Nikolat27/king_of_numbers/backend/utils"
)

func (handler *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, 200, "Backend is working!")
}
