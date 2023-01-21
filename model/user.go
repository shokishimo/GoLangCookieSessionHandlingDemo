package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shokishimo/OneTap/db"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SaveUser(user User) error {
	// get keys to access database
	if err := godotenv.Load(); err != nil {
		return err
	}
	database := os.Getenv("DATABASE")
	userCollection := os.Getenv("COLLECTION_USER")
	if database == "" || userCollection == "" {
		fmt.Println(database)
		fmt.Println(userCollection)
		return errors.New("database access key strings are wrong")
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
