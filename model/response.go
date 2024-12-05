package model

type Response struct {
	Message string
}

// orders
type OrderCreateSectionRq struct {
	IdTable string `json:"id_table"`
}

type GetOrdersRq struct {
	IdTable string `json:"id_table"`
}

type PayloadOrderItens struct {
	IdOrder string `json:"id_order"`
}
