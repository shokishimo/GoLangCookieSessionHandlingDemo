package rooter

import (
	"html/template"
	"net/http"
)

func CreateNewProjectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("static/public/createNewProject.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	tmpl.Execute(w, nil)
}
