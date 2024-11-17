package model

type Product struct {
	ID       int     `json:"id_product"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Volume   float32 `json:"volume"`
	Describe string  `json:"describe"`
}
