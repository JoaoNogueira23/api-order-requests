package model

import "time"

type Section struct {
	IdSection string    `json:"id_section"`
	IdTable   int       `json:"id_table"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"`
}

type Order struct {
	IdOrder   string    `json:"id_order"`
	IdSection string    `json:"id_section"`
	OrderTime time.Time `json:"order_time"`
	Status    string    `json:"status"`
}

type OrderItem struct {
	IdOrderItem string  `json:"id_order_item"`
	IdOrder     string  `json:"id_order"`
	IdProduct   string  `json:"id_product"`
	Quantity    int     `json:"quantity"`
	UnityPrice  float32 `json:"unit_price"`
	TotalPrice  float32 `json:"total_price"`
}
