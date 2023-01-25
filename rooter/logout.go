package rooter

import (
	"fmt"
	user "github.com/shokishimo/OneTap/model"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// delete cookie from both database and browser
	cookie, _ := r.Cookie("sessionID")
	// obtain sessionID inside the cookie
	sessionID := cookie.Value
	err := user.DeleteCookie(w, sessionID)
	if err != nil {
		fmt.Fprint(w, "log out system error")
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
