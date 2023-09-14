package models

type BaseModel struct {
	Name    string    `json:"name"`
	Table   Table     `json:"table"`
	Columns []Columns `json:"columns"`
}

type Table struct {
	Name    string   `json:"name"`
	Joins   []Joins  `json:"joins"`
	Deleted *Deleted `json:"deleted"`
}

type Columns struct {
	Field string `json:"field"`
	Label string `json:"label"`
}

type Joins struct {
	Name    string `json:"name"`
	Foreign string `json:"foreign"`
	Key     string `json:"key"`
	Join    string `json:"join"`
}

type Deleted struct {
	Field string `json:"field"`
	Value string `json:"value"`
}
