package rooter

import (
	"html/template"
	"net/http"
)

func AccountHomeHandler(w http.ResponseWriter, r *http.Request) {
	// for those who don't sign up yet
	tmpl, err := template.ParseFiles("static/public/accountHome.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	tmpl.Execute(w, nil)
}
