package models

type BaseModel struct {
	Name    string    `json:"name"`
	Table   Table     `json:"table"`
	Columns []Columns `json:"columns"`
}

type Table struct {
	Name string `json:"name"`
}

type Columns struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}
