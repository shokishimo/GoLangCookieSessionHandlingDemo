package rooter

import (
	"fmt"
	user "github.com/shokishimo/OneTap/model"
	"html/template"
	"net/http"
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// get sessionID
	cookie, _ := r.Cookie("sessionID")
	sessionID := cookie.Value
	// get user data with sessionID
	theUser, result := user.GetUserBySessionID(sessionID)
	if !result {
		fmt.Fprint(w, "Unable to get user data")
		return
	}

	// for those who don't sign up yet
	tmpl, err := template.ParseFiles("static/public/account.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	tmpl.Execute(w, theUser)
	return
}
