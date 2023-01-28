package rooter

import (
	"net/http"
	"os/exec"
)

func DemoStartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cmd := exec.Command("open", "-a", "/Applications/Postman.app")
	err := cmd.Start()
	if err != nil {
		// handle error
	}
	cmd = exec.Command("open", "-a", "Google Chrome", "https://github.com")
	err = cmd.Start()
	if err != nil {
		// handle error
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
