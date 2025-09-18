package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Nikolat27/king_of_numbers/backend/types"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encodedData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		WriteErrorJSON(w, http.StatusInternalServerError, types.NewErrorResponse("encodingData", err))
		return
	}

	if _, err := w.Write(encodedData); err != nil {
		WriteErrorJSON(w, http.StatusInternalServerError, types.NewErrorResponse("writingJSON", err))
	}
}

func ParseJSON(w http.ResponseWriter, requestBody io.ReadCloser, maxBytes int64, input any) {
	body, err := io.ReadAll(io.LimitReader(requestBody, maxBytes))
	if err != nil {
		WriteErrorJSON(w, http.StatusInternalServerError, types.NewErrorResponse("readJSON", err))
		return
	}

	if err := json.Unmarshal(body, input); err != nil {
		WriteErrorJSON(w, http.StatusInternalServerError, types.NewErrorResponse("parseJSON", err))
	}
}

func WriteErrorJSON(w http.ResponseWriter, statusCode int, response *types.ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := map[string]*types.ErrorResponse{
		"error": response,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
