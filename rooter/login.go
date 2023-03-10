package rooter

import (
	"context"
	"errors"
	"fmt"
	"github.com/shokishimo/OneTap/db"
	user "github.com/shokishimo/OneTap/model"
	"go.mongodb.org/mongo-driver/bson"
	"html/template"
	"net/http"
)

type loginError struct {
	DidFail bool
	Fail    string
}

// LoginHandler handles both get and post method
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
	if r.Method == http.MethodGet { // handle GET method
		res := r.URL.Query().Get("error")
		if res == "invalid_credentials" {
			data := loginError{DidFail: true, Fail: "Invalid username or password"}
			loginGet(w, true, data)
		} else {
			loginGet(w, false, loginError{})
		}
	} else if r.Method == http.MethodPost { // handle POST method
		err := loginPost(w, r)
		if err != nil {
			fmt.Println("login failed")
			// Redirect to login page again
			http.Redirect(w, r, "/login?error=invalid_credentials", http.StatusSeeOther)
		}
	} else { // others
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

// loginGet returns a login html page
func loginGet(w http.ResponseWriter, didFail bool, data loginError) {
	tmpl, err := template.ParseFiles("static/public/login.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if !didFail {
		tmpl.Execute(w, nil)
	} else {
		tmpl.Execute(w, data)
	}
}

// loginPost handles login
func loginPost(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := user.Hash(r.FormValue("password"))

	// TODO: validate the user input

	// get access keys
	database, userCollection, err := user.GetDatabaseAccessKeysForUsers()
	if err != nil {
		return err
	}
	// database connection
	client, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Disconnect(client)
	collection := client.Database(database).Collection(userCollection)

	// create a new sessionID
	sessionID := user.GenerateSessionID()
	// update the user in the database if the user exists
	filter := bson.M{"username": username, "password": password}
	update := bson.M{"$set": bson.M{"sessionid": user.Hash(sessionID)}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)

	// when an error happened in the transaction
	if err != nil {
		fmt.Println("An error occurred when login check to the database")
		return err
	}
	// when the user with the username and password not found
	if result.MatchedCount == 0 {
		fmt.Println("user not found")
		return errors.New("user not found")
	}

	// when the user is found
	// save the cookie in the client browser
	user.SetCookie(w, sessionID)

	// Redirect to account home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}
