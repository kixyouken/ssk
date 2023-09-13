package models

type BaseModel struct {
	Name    string    `json:"name"`
	Table   Table     `json:"table"`
	Columns []Columns `json:"columns"`
}

type Table struct {
	Name  string  `json:"name"`
	Joins []Joins `json:"joins"`
}

type Columns struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

type Joins struct {
	Name    string `json:"name"`
	Foreign string `json:"foreign"`
	Key     string `json:"key"`
	Join    string `json:"join"`
}
