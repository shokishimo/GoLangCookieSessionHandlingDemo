package rooter

import (
	"context"
	"fmt"
	"github.com/shokishimo/OneTap/db"
	user "github.com/shokishimo/OneTap/model"
	"go.mongodb.org/mongo-driver/bson"
	"html/template"
	"net/http"
)

// LoginHandler handles both get and post method
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet { // handle GET method
		loginGet(w)
	} else if r.Method == http.MethodPost { // handle POST method
		err := loginPost(w, r)
		if err != nil {
			fmt.Println("login failed")
			// Redirect to account home page
			http.Redirect(w, r, "/loginFail", http.StatusSeeOther)
		}
	} else { // others
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

// loginGet returns a login html page
func loginGet(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("static/public/login.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

// loginPost handles login
func loginPost(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := user.Hash(r.FormValue("password"))

	// TODO: validate the user input

	// get access keys
	database, userCollection, err := user.GetDatabaseAccessKeys()
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
	_, err = collection.UpdateOne(context.TODO(), filter, update)

	// when the user with the username and password not found
	if err != nil {
		fmt.Println("User not found")
		return err
	}

	// save the cookie in the client browser
	user.SetCookie(w, sessionID)

	// Redirect to account home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}
