package model

type Table struct {
	ID           string `json:"id_table"`
	Table_number string `json:"table_number"`
	Status       string `json:"status"`
	Location     string `json:"location"`
	Capacity     int    `json:"capacity"`
}
