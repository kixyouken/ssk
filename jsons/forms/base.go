package forms

type BaseForm struct {
	Name   string `json:"name"`
	Action Action `json:"action"`
}

type Action struct {
	Bind Bind `json:"bind"`
}

type Bind struct {
	Model string `json:"model"`
}
