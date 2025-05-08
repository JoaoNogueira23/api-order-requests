package model

import (
	"database/sql"
)

type Products struct {
	ID          string         `json:"id_product"`
	Name        string         `json:"name"`
	Price       float32        `json:"price"`
	Volume      float32        `json:"volume"`
	Describe    string         `json:"description"`
	Isactive    bool           `json:"isacticve"`
	Ispromotion bool           `json:"ispromotion"`
	Discount    float32        `json:"discount"`
	UrlImage    sql.NullString `json:"url_image"`
}

type ProductsRequestOrder struct {
	ID       string  `json:"id_product"`
	Price    float32 `json:"price"`
	Volume   float32 `json:"volume"`
	Quantity int     `json:"quantity"`
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
