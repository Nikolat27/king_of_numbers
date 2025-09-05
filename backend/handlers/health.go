package handlers

import (
	"net/http"
)

func (handler *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Backend is working!"))
}
