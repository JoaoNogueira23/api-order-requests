package model

type Response struct {
	Message string
}

// orders
type OrderCreateSectionRq struct {
	IdTable string `json:"id_table"`
}

type PayloadGetOrdersRq struct {
	IdTable string `json:"id_table"`
}

type PayloadOrderItens struct {
	IdOrder string `json:"id_order"`
}
