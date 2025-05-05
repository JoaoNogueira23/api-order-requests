package model

type Products struct {
	ID          string  `json:"id_product"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Volume      float32 `json:"volume"`
	Describe    string  `json:"description"`
	Isactive    bool    `json:"isacticve"`
	Ispromotion bool    `json:"ispromotion"`
	Discount    float32 `json:"discount"`
	UrlImage    string  `json:"url_image"`
}

type ProductsRequestOrder struct {
	ID         int     `json:"id_product"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	Volume     float32 `json:"volume"`
	Quantity   int     `json:"quantity"`
	TotalPrice float32 `json:"total_price"`
}

// model to requests new products on table
type OrderCreateRq struct {
	ProductsList []ProductsRequestOrder `json:"products_list"`
	IdSection    string                 `json:"id_section"`
}

type ProductResponse struct {
	Page    string      `json:"page"`
	PerPage string      `json:"per_page"`
	Total   int         `json:"total"`
	Data    interface{} `json:"data"`
}
