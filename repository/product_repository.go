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

func (pr *ProductRepository) GetProducts() ([]model.Products, error) {

	query := "SELECT id_product, name, price, description, volume, isactive, ispromotion, discount FROM products"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Products{}, err
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
			&productObj.Discount)

		if err != nil {
			fmt.Println(err)
			return []model.Products{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
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

func (pr *ProductRepository) GetProductById(id_product int) (*model.Products, error) {

	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Products

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
		&produto.Volume,
		&produto.Describe,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &produto, nil
}
