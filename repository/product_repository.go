package repository

import (
	"api-blog-go/model"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts(page int, limit int) ([]model.Products, int, error) {
	query := fmt.Sprintf(`
		SELECT 
			id_product, 
			name, 
			price, 
			description, 
			volume, 
			isactive, 
			ispromotion, 
			discount,
			url_image
		FROM products
		limit %d offset %d
	`, limit, limit*page)

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Products{}, -1, err
	}

	var productList []model.Products
	var productObj model.Products

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
			&productObj.Describe,
			&productObj.Volume,
			&productObj.Isactive,
			&productObj.Ispromotion,
			&productObj.Discount,
			&productObj.UrlImage)

		if err != nil {
			fmt.Println(err)
			return []model.Products{}, -1, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	// data total storage
	// Executar a query
	var count int
	err = pr.connection.QueryRow("SELECT COUNT(*) FROM Products").Scan(&count)
	if err != nil {
		return []model.Products{}, -1, err
	}

	return productList, count, nil
}

func (pr *ProductRepository) CreateProduct(product model.Products) (string, error) {

	var id string

	// variables
	// Define o seed para o gerador de números aleatórios
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Gera um ULID (ordenado lexicograficamente)
	id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	query, err := pr.connection.Prepare("INSERT INTO products" +
		"(id_product,name, price, volume, description)" +
		" VALUES ($1, $2, $3, $4, $5) RETURNING id_product")

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = query.QueryRow(id, product.Name, product.Price, product.Volume, product.Describe).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(id_product string) (*model.Products, error) {

	query, err := pr.connection.Prepare("SELECT * FROM products WHERE id_product = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Products

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Describe,
		&produto.Price,
		&produto.Volume,
		&produto.Isactive,
		&produto.Ispromotion,
		&produto.Discount)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &produto, nil
}
