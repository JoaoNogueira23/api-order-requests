package model

type Products struct {
	ID          int     `json:"id_product"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Volume      float32 `json:"volume"`
	Describe    string  `json:"description"`
	Isactive    bool    `json:"isacticve"`
	Ispromotion bool    `json:"ispromotion"`
	Discount    float32 `json:"discount"`
}

type ProductsRequestOrder struct {
	ID         int     `json:"id_product"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	Volume     float32 `json:"volume"`
	Quantity   int     `json:"quantity"`
	TotalPrice float32 `json:"total_price"`
}
