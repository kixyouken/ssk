package tables

type BaseTable struct {
	Name   string `json:"name"`
	Action Action `json:"action"`
}

type Action struct {
	Bind   Bind     `json:"bind"`
	Wheres []Wheres `json:"wheres"`
	Orders []Orders `json:"orders"`
	Page   int      `json:"page"`
	Limit  int      `json:"limit"`
	Count  int      `json:"count"`
}

type Bind struct {
	Model string `json:"model"`
}

type Wheres struct {
	Field  string `json:"field"`
	Search string `json:"search"`
	Type   string `json:"type"`
}

type Orders struct {
	Field string `json:"field"`
	Sort  string `json:"sort"`
}
