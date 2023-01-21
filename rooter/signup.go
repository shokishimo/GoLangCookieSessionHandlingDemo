package rooter

import (
	"fmt"
	"github.com/shokishimo/OneTap/model"
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
	user := model.User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	// TODO: Add validations to input username and password

	// TODO: Hash the password

	// save the user
	err := model.SaveUser(user)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	// TODO: Create a cookie session for the user

	// success log
	fmt.Fprint(w, "successfully inserted the user")

	// Redirect to account home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
