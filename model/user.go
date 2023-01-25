package model

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shokishimo/OneTap/db"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	SessionID string `json:"sessionid"`
}

// SaveUser saves user data into a database
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
		Expires:  time.Now().Add(3600 * 24 * 3 * time.Second), // 3 days
		HttpOnly: true,
		Secure:   false, // When developing
		//Secure:   true,               // When deploying
		SameSite: http.SameSiteLaxMode, // TODO: change this to Strict (maybe)
	}
	http.SetCookie(w, &cookie)
}

// DeleteCookie deletes cookie from both the database and browser
func DeleteCookie(w http.ResponseWriter, sid string) error {
	err := DeleteCookieFromDatabase(sid)
	if err != nil {
		return err
	}

	// delete cookie from browser
	cookie := &http.Cookie{
		Name:     "sessionID",
		Value:    "",
		Expires:  time.Now(),
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	return nil
}

func DeleteCookieFromDatabase(sid string) error {
	// Hash the sid
	hashed := Hash(sid)
	// get access keys
	database, userCollection, err := GetDatabaseAccessKeys()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// database connection
	client, err := db.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Disconnect(client)
	collection := client.Database(database).Collection(userCollection)

	filter := bson.M{"sessionid": hashed}
	update := bson.M{"$set": bson.M{"sessionid": ""}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)

	// when an error happened in the transaction
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// when the user with the username and password not found
	if result.MatchedCount == 0 {
		fmt.Println("user not found")
		return errors.New("user not found")
	}

	fmt.Println("Delete cookie from database")
	return nil
}
