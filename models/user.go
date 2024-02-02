package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//The user entity
type User struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Username string `json:"username"`

	Email string `json:"email"`

	Password string `json:"password"`
}
