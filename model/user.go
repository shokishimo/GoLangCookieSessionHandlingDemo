package model

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	"github.com/shokishimo/OneTap/db"
	"os"
)

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	SessionID string `json:"session_id"`
}

func SaveUser(user User) error {
	// get access keys
	database, userCollection, err := getDatabaseAccessKeys()
	if err != nil {
		return err
	}

	// connect to database
	client, err := db.Connect()
	if err != nil {
		return err
	}

	// begin insert user
	collection := client.Database(database).Collection(userCollection)
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	// Disconnect from db
	defer db.Disconnect(client)
	return nil
}

// getDatabaseAccessKeys gets keys to access database from .env file
func getDatabaseAccessKeys() (database string, userCollection string, err error) {
	if err := godotenv.Load(); err != nil {
		return "", "", err
	}
	database = os.Getenv("DATABASE")
	userCollection = os.Getenv("COLLECTION_USER")
	if database == "" || userCollection == "" {
		return "", "", errors.New("database access key strings are wrong")
	}
	return database, userCollection, nil
}