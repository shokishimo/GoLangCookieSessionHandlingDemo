package rooter

import (
	"fmt"
	"github.com/shokishimo/OneTap/model"
	"html/template"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		signUpGet(w)
	} else if r.Method == http.MethodPost {
		signUpPost(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

func signUpGet(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("static/public/signup.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func signUpPost(w http.ResponseWriter, r *http.Request) {
	user := model.User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	// For debugging
	// fmt.Fprintf(w, "Username: %s, Password: %s", user.Username, user.Password)

	err := model.SaveUser(user)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, "successfully inserted the user")
}
