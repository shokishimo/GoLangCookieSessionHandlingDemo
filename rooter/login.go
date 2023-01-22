package rooter

import (
	user "github.com/shokishimo/OneTap/model"
	"html/template"
	"net/http"
)

// LoginHandler handles both get and post method
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet { // handle GET method
		loginGet(w)
	} else if r.Method == http.MethodPost { // handle POST method
		loginPost(w, r)
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
func loginPost(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := user.Hash(r.FormValue("password"))

	// TODO: validate the user input

	// check if the input user already exists in the database

	sessionID := user.GenerateSessionID()
	// save the cookie in the client browser
	user.SetCookie(w, sessionID)

	// Redirect to account home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
