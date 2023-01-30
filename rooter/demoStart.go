package rooter

import (
	"fmt"
	app "github.com/shokishimo/OneTap/model"
	"net/http"
	"os/exec"
)

func DemoStartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("'/demoStart' is accessed")
	cmd := exec.Command(app.MacosCommand1, app.MacosCommand2, "/Applications/Postman.app")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}

	cmd = exec.Command(app.MacosCommand1, app.MacosCommand2, "Google Chrome", "https://github.com")
	err = cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
