package dto

import (
	"encoding/json"
	"net/http"
)

// The error response struct
type ErrorResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

type LoginRes struct {
	Message      string `json:"message"`
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

// The response function
func SendResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	response := Response{Message: message}
	json.NewEncoder(w).Encode(response)
}

// The error function
func SendErrorResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	response := ErrorResponse{Message: message}
	json.NewEncoder(w).Encode(response)
}

func LoginResponse(w http.ResponseWriter, status int, message string, access string, refresh string) {
	w.WriteHeader(status)
	accessToken := access
	refreshToken := refresh
	response := LoginRes{
		Message:      message,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	json.NewEncoder(w).Encode(response)
}
