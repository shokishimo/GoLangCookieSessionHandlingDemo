package rooter

import (
	app "github.com/shokishimo/OneTap/model"
	"net/http"
	"os/exec"
)

func DemoStartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cmd := exec.Command(app.MacOS_Command1, app.MacOS_Command2, "/Applications/Postman.app")
	err := cmd.Start()
	if err != nil {
		// handle error
	}
	cmd = exec.Command(app.MacOS_Command1, app.MacOS_Command2, "Google Chrome", "https://github.com")
	err = cmd.Start()
	if err != nil {
		// handle error
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
