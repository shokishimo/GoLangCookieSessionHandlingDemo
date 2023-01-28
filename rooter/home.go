package rooter

import (
	"fmt"
	user "github.com/shokishimo/OneTap/model"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("'/' is accessed")
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
	theUser, result := user.GetUserBySessionID(sessionID)
	// for those who did signup
	if result {
		fmt.Println(theUser)
		// Redirect to account home page
		http.Redirect(w, r, "/accountHome", http.StatusSeeOther)
		return
	} else {
		ShowPublicHome(w)
		return
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
