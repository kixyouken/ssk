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
	Model  string `json:"model"`
	Filter Filter `json:"filter"`
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

type Filter struct {
	Distinct []Distinct `json:"distinct"`
}

type Distinct struct {
	Field string `json:"field"`
}
