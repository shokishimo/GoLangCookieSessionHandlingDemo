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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// obtain cookie
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		fmt.Println("cookie doesn't exist")
		// when there is no cookie set to the browser
		ShowPublicHome(w)
		return
	}
	// obtain sessionID inside the cookie
	sessionID := cookie.Value

	// check if the sessionID exists, if so,
	theUser, result := doesSessionIDExist(sessionID)

	// for those who did signup
	if result {
		fmt.Fprintf(w, "Username: "+theUser.Username+
			", password: "+theUser.Password+
			", sessionID: "+theUser.SessionID)
	} else {
		ShowPublicHome(w)
	}
}

// ShowPublicHome shows the public template home to the browser when the user doesn't sign in
func ShowPublicHome(w http.ResponseWriter) {
	// for those who don't sign up yet
	tmpl, err := template.ParseFiles("static/public/home.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	tmpl.Execute(w, nil)
}

// isSessionIDValid checks if the sessionID exists. It return user.User and true if the user with the sessionID exists
func doesSessionIDExist(sid string) (user.User, bool) {
	// Hash the sid
	hashed := user.Hash(sid)
	// get access keys
	database, userCollection, err := user.GetDatabaseAccessKeys()
	if err != nil {
		fmt.Println(err.Error())
		return user.User{}, false
	}
	// database connection
	client, err := db.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return user.User{}, false
	}
	defer db.Disconnect(client)
	collection := client.Database(database).Collection(userCollection)

	// bson.D creates a set of key and value to filter the database
	// bson.M creates a map, bson.A creates an array
	var result bson.M
	// Define the filter to find a specific document
	filter := bson.M{"session_id": hashed}
	// check if the sessionID exists in the database
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	// when the user with the sessionID not found
	if err != nil {
		fmt.Println(err.Error())
		return user.User{}, false
	}
	// when the user is found
	fmt.Println("user found")
	return user.User{
		Username:  result["username"].(string),
		Password:  result["password"].(string),
		SessionID: result["session_id"].(string),
	}, true
}
