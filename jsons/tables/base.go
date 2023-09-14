package tables

type BaseTable struct {
	Name   string `json:"name"`
	Action Action `json:"action"`
}

type Action struct {
	Bind   Bind     `json:"bind"`
	Wheres []Wheres `json:"wheres"`
}

type Bind struct {
	Model string `json:"model"`
}

type Wheres struct {
	Field  string `json:"field"`
	Search string `json:"search"`
	Type   string `json:"type"`
}
