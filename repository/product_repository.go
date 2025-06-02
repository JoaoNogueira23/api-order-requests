package repository

import (
	"api-blog-go/model"
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
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
			url_image,
			category
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
			&productObj.UrlImage,
			&productObj.Category)

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

func (pr *ProductRepository) CreateProduct(products []model.Products) (*string, error) {
	// variables
	var (
		values       []interface{}
		placeholders []string
	)

	// Define o seed para o gerador de números aleatórios
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i, products := range products {
		// Gera um ULID (ordenado lexicograficamente)
		id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

		start := i*5 + 1

		placeholders = append(placeholders,
			fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", start, start+1, start+2, start+3, start+4, start+5))

		values = append(values,
			id,
			products.Name,
			products.Price,
			products.Volume,
			products.Describe,
			products.Category,
		)
	}

	query := fmt.Sprintf(`
		INSERT INTO products (id_product, name, price, volume, description)
		VALUES %s`, strings.Join(placeholders, ","))

	_, err := pr.connection.Exec(query, values...)

	if err != nil {
		return nil, fmt.Errorf("erro ao inserir produtos: %w", err)
	}

	message := "Produto(s) criados com sucesso!"

	return &message, nil
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
		&produto.Discount,
		&produto.UrlImage)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	query.Close()
	return &produto, nil
}
