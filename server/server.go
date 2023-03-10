package server

import (
	"github.com/shokishimo/OneTap/rooter"
	"net/http"
)

// ServeMux creates a new HTTP server.
func ServeMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rooter.HomeHandler)
	mux.HandleFunc("/signup", rooter.SignUpHandler)
	mux.HandleFunc("/login", rooter.LoginHandler)
	mux.HandleFunc("/accountHome", rooter.AccountHomeHandler)
	mux.HandleFunc("/logout", rooter.LogoutHandler)
	mux.HandleFunc("/account", rooter.AccountHandler)
	mux.HandleFunc("/createNewProject", rooter.CreateNewProjectHandler)

	mux.HandleFunc("/demoStart", rooter.DemoStartHandler)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	return mux
}
