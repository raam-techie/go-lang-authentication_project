package main

import (
	"authservice/db"
	"authservice/handlers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// The main function
func main() {

	db.InitMongoDB()

	router := httprouter.New()

	router.POST("/api/signup", handlers.CreateUserHandler)
	router.POST("/api/login", handlers.Loginhandler)
	err := http.ListenAndServe(":8000", router)

	if err != nil {
		log.Fatal("Localhost not connected")
	}

}
