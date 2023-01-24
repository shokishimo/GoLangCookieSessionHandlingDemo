package rooter

import (
	"fmt"
	"html/template"
	"net/http"
)

func AccountHomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// check if the cookie is valid
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		fmt.Println("cookie doesn't exist")
		// when there is no cookie set to the browser
		ShowPublicHome(w)
		return
	}
	sessionID := cookie.Value

	// check if the sessionID exists, if so,
	_, result := doesSessionIDExist(sessionID)
	// for those who did signup
	if result {
		// for those who don't sign up yet
		tmpl, err := template.ParseFiles("static/public/accountHome.html")
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		tmpl.Execute(w, nil)
		return
	} else {
		ShowPublicHome(w)
		return
	}
}
