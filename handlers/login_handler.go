package handlers

import (
	"authservice/db"
	"authservice/dto"
	model "authservice/models"
	"authservice/security"
	"authservice/utility"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Loginhandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	login(w, r)
}

func login(w http.ResponseWriter, r *http.Request) {

	var credentials model.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Query the user from the database based on username and password
	user, err := db.GetUserByEmail(credentials.Email)

	if user.Email == "" {
		dto.SendErrorResponse(w, http.StatusNotFound, "User Not Found")
		return
	}

	if !utility.CheckPasswordHash(credentials.Password, user.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	accesstoken, err := security.GenerateJWT(user.Username)
	fmt.Println(accesstoken)
	refreshToken, err := security.GenerateRefreshToken()
	fmt.Println(refreshToken)
	fmt.Printf("Users: %v\n", user)
	dto.LoginResponse(w, http.StatusOK, "Credentail accepted", accesstoken, refreshToken)
}
