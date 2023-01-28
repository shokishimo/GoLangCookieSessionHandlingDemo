package model

type Model struct {
	Username    string `json:"username"`
	ProjectName string `json:"projectname"`
	Os          string `json:"os"`
	Browser     string `json:"browser"`
	Url         string `json:"url"`
}

const (
	MacOS_Command1 string = "open"
	MacOS_Command2 string = "-a"

	Windows_Command string = "start"
)
