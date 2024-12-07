package model

type Table struct {
	ID           string `json:"id_table"`
	Table_number string `json:"table_number"`
	IsOccupied   bool   `json:"isOccupied"`
	Location     string `json:"location"`
}
