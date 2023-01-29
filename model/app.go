package model

type App struct {
	Username    string `json:"username"`
	ProjectName string `json:"projectname"`
	Os          string `json:"os"`
	Browser     string `json:"browser"`
	Url         string `json:"url"`
}

const (
	MacosCommand1 string = "open"
	MacosCommand2 string = "-a"

	WindowsCommand string = "start"
)
