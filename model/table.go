package model

type Table struct {
	ID           int    `json:"id_table"`
	Table_number int    `json:"table_number"`
	IsOccupied   bool   `json:"isOccupied"`
	Location     string `json:"location"`
}
