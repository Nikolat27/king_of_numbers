package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encodedData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(encodedData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ParseJSON(w http.ResponseWriter, requestBody io.ReadCloser, maxBytes int64, input any) {
	body, err := io.ReadAll(io.LimitReader(requestBody, maxBytes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WriteErrorJSON(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]string{
		"error": message,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
