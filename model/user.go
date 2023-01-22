package model

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/joho/godotenv"
	"github.com/shokishimo/OneTap/db"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	SessionID string `json:"sessionid"`
}

func SaveUser(user User) error {
	// get access keys
	database, userCollection, err := GetDatabaseAccessKeys()
	if err != nil {
		return err
	}
	// connect to database
	client, err := db.Connect()
	if err != nil {
		return err
	}
	// Disconnect from db
	defer db.Disconnect(client)

	// begin insert user
	collection := client.Database(database).Collection(userCollection)
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

// GetDatabaseAccessKeys gets keys to access database from .env file
func GetDatabaseAccessKeys() (database string, userCollection string, err error) {
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

// GenerateSessionID Generates a sessionID
func GenerateSessionID() string {
	b := make([]byte, 16)
	rand.Read(b)
	sessionID := hex.EncodeToString(b)
	return sessionID
}

// Hash hashes the input and return the hashed value
func Hash(val string) string {
	// Hash the password
	hash := sha256.Sum256([]byte(val))
	// Encode the hash as a hexadecimal string
	hashed := hex.EncodeToString(hash[:])
	return hashed
}

// SetCookie sets a cookie
func SetCookie(w http.ResponseWriter, sid string) {
	cookie := http.Cookie{
		Name:     "sessionID",
		Value:    sid,
		Expires:  time.Now().Add(300 * 1 * time.Second),
		HttpOnly: true,
		Secure:   false,                // TODO: change this to true when deploying
		SameSite: http.SameSiteLaxMode, // TODO: change this to Strict (maybe)
	}
	http.SetCookie(w, &cookie)
}
