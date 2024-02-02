package db

import (
	model "authservice/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client         *mongo.Client
	db             *mongo.Database
	err            error
	userCollection *mongo.Collection
)

func InitMongoDB() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")

	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB created!!!")
	db = client.Database("user_authentication")

	userCollection = db.Collection("users")
}

func InsertUser(user model.User) error {
	_, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (model.User, error) {
	var user model.User

	// Query the user based on email
	filter := bson.M{"email": email}
	err := userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
