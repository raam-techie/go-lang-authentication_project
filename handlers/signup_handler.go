package handlers

import (
	"authservice/db"
	"authservice/dto"
	model "authservice/models"
	"authservice/utility"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	create(w, r)
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	//Hashing password
	pass, err := utility.HashPassword(user.Password)
	if err != nil {
		fmt.Println("hash function error")
	}
	user.Password = pass
	err = db.InsertUser(user)
	if err != nil {
		dto.SendErrorResponse(w, http.StatusInternalServerError, "Failed to save user")
	}

	dto.SendResponse(w, http.StatusCreated, "User created successfully")
	fmt.Printf("Users: %v\n", user)
}
