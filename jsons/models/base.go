package models

type BaseModel struct {
	Name    string    `json:"name"`
	Table   Table     `json:"table"`
	Columns []Columns `json:"columns"`
}

type Table struct {
	Name       string       `json:"name"`
	Joins      []Joins      `json:"joins"`
	Withs      []Withs      `json:"withs"`
	WithsCount []WithsCount `json:"Withs_count"`
	Deleted    *Deleted     `json:"deleted"`
}

type Columns struct {
	Field string  `json:"field"`
	Label string  `json:"label"`
	Attrs []Attrs `json:"attr"`
}

type Attrs struct {
	In  string `json:"in"`
	Out string `json:"out"`
}

type Joins struct {
	Name    string    `json:"name"`
	Foreign string    `json:"foreign"`
	Key     string    `json:"key"`
	Join    string    `json:"join"`
	Wheres  []Wheres  `json:"wheres"`
	Columns []Columns `json:"columns"`
}

type Deleted struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type Wheres struct {
	Field  string `json:"field"`
	Search string `json:"search"`
	Type   string `json:"type"`
	Value  string `json:"value"`
}

type Withs struct {
	Name    string    `json:"name"`
	Foreign string    `json:"foreign"`
	Key     string    `json:"key"`
	Join    string    `json:"join"`
	Wheres  []Wheres  `json:"wheres"`
	Columns []Columns `json:"columns"`
	Orders  []Orders  `json:"orders"`
}

type Orders struct {
	Field string `json:"field"`
	Sort  string `json:"sort"`
}

type WithsCount struct {
	Name    string   `json:"name"`
	Foreign string   `json:"foreign"`
	Key     string   `json:"key"`
	Join    string   `json:"join"`
	Wheres  []Wheres `json:"wheres"`
}
