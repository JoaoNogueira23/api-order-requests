package model

type Response struct {
	Message string
}

// orders
type OrderCreateSectionRq struct {
	IdTable int `json:"id_table"`
}

type GetOrdersRq struct {
	IdTable int `json:"id_table"`
}
