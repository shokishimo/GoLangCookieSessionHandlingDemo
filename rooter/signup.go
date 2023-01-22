package rooter

import (
	"fmt"
	user "github.com/shokishimo/OneTap/model"
	"html/template"
	"net/http"
)

// SignUpHandler handles the api endpoint '/signup'
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet { // handle GET method
		signUpGet(w)
	} else if r.Method == http.MethodPost { // handle POST method
		signUpPost(w, r)
	} else { // others
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

// signUpGet return signup.html file
func signUpGet(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("static/public/signup.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

// signUpPost save a user signed up
func signUpPost(w http.ResponseWriter, r *http.Request) {
	sessionID := user.GenerateSessionID()
	theUser := user.User{
		Username:  r.FormValue("username"),
		Password:  user.Hash(r.FormValue("password")),
		SessionID: user.Hash(sessionID),
	}
	// TODO: validate the user input

	// save the user
	err := user.SaveUser(theUser)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	// success log
	fmt.Println("successfully inserted the user")

	// save the cookie in the client browser
	user.SetCookie(w, sessionID)

	// Redirect to account home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
