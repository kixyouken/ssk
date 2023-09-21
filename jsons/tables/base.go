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
	Model     string     `json:"model"`
	Filter    Filter     `json:"filter"`
	Paginate  *bool      `json:"paginate"`
	Recursion *Recursion `json:"recursion"`
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

type Recursion struct {
	ParentID string `json:"parent_id"`
	ChildID  string `json:"child_id"`
}
