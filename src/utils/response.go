package utils

import (
	"encoding/json"
	"go31ngo/src/models"
	"net/http"
)

func SendApiResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode == http.StatusNoContent {
		return
	}

	if statusCode >= 400 {
		json.NewEncoder(w).Encode(models.ApiResponse{
			Success: false,
			Message: data.(string),
		})
	} else {
		json.NewEncoder(w).Encode(models.ApiResponse{
			Success: true,
			Data:    data,
		})
	}
}
