package tables

type BaseTable struct {
	Name   string `json:"name"`
	Action Action `json:"action"`
}

type Action struct {
	Bind   Bind   `json:"bind"`
	Search Search `json:"search"`
}

type Bind struct {
	Table string `json:"table"`
}

type Search struct {
}
