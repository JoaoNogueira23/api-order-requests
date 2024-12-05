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

type OrderRepository struct {
	conn *sql.DB
}

func NewOrderRepository(conn *sql.DB) OrderRepository {
	return OrderRepository{
		conn: conn,
	}
}

func (or *OrderRepository) CreateSection(id_table string) (string, error) {
	var id string
	location, err := time.LoadLocation("America/Sao_Paulo")
	// Gera o datetime atual na região de São Paulo
	now := time.Now().In(location)

	query, err := or.conn.Prepare("INSERT INTO sections" +
		"(id_section, id_table, start_time)" +
		"VALUES ($1, $2, $3) RETURNING id_section;")

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// variables
	// Define o seed para o gerador de números aleatórios
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Gera um ULID (ordenado lexicograficamente)
	id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	err = query.QueryRow(id, id_table, now).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	query.Close()
	return id_table, nil
}

func (or *OrderRepository) CreateOrder(id_section string) (string, error) {
	var id string
	location, err := time.LoadLocation("America/Sao_Paulo")
	// Gera o datetime atual na região de São Paulo
	now := time.Now().In(location)

	query, err := or.conn.Prepare("INSERT INTO orders" +
		"(id_order, id_section, order_time)" +
		"VALUES ($1, $2, $3) RETURNING id_order;")

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// variables
	// Define o seed para o gerador de números aleatórios
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Gera um ULID (ordenado lexicograficamente)
	id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	err = query.QueryRow(id, id_section, now).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	query.Close()
	return id, nil
}

func (or *OrderRepository) CreateOrderItem(id_order string, productsList []model.ProductsRequestOrder) (int, error) {
	var builder strings.Builder

	// Query base
	builder.WriteString("INSERT INTO orderitems (id_order_item, id_order, id_product, quantity, unit_price, total_price) VALUES ")

	for i, product := range productsList {
		// Gera o ULID único
		entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
		id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

		// Adiciona os valores ao builder
		builder.WriteString(fmt.Sprintf(
			"('%s', '%s', '%s', %d, %.2f, %.2f)",
			id, id_order, product.ID, product.Quantity, product.Price, product.TotalPrice,
		))

		// Adiciona vírgula entre valores, exceto no último item
		if i < len(productsList)-1 {
			builder.WriteString(", ")
		}
	}

	builder.WriteString(";") // Fecha a query

	result, err := or.conn.Exec(builder.String())

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	rowsEffected, err := result.RowsAffected()

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return int(rowsEffected), nil
}

func (or *OrderRepository) GetOrders(id_table int) ([]model.Order, error) {
	query := `
		SELECT 
			id_order, 
			ORD.id_section, 
			order_time,
			ORD.status
		FROM sections SEC
		LEFT JOIN orders ORD
		ON SEC.id_section = ORD.id_section
		WHERE SEC.id_table = $1`
	rows, err := or.conn.Query(query, id_table)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var ordersList []model.Order
	var orderObj model.Order

	if !rows.Next() {
		return nil, fmt.Errorf("Section not found!")
	}

	for rows.Next() {
		err = rows.Scan(
			&orderObj.IdOrder,
			&orderObj.IdSection,
			&orderObj.OrderTime,
			&orderObj.Status)

		if err != nil {
			return nil, err
		}

		ordersList = append(ordersList, orderObj)
	}

	return ordersList, nil
}

func (or *OrderRepository) GetOrderItens(id_order string) ([]model.OrderItemRq, error) {
	query := `SELECT
			SUM(total_price) as total_price,
			SUM(quantity) as quantity,
			PRT.name as product_name
		FROM ORDERITEMS ORDI
		LEFT JOIN PRODUCTS PRT
		ON PRT.id_product = ORDI.id_product
		WHERE id_order = $1
		GROUP BY
			ORDI.id_product,
			PRT.name`
	rows, err := or.conn.Query(query, id_order)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var ordersList []model.OrderItemRq
	var orderObj model.OrderItemRq

	if !rows.Next() {
		return nil, fmt.Errorf("Section not found!")
	}

	for rows.Next() {
		err = rows.Scan(
			&orderObj.Total_price,
			&orderObj.Quantity,
			&orderObj.ProductName)

		if err != nil {
			return nil, err
		}

		ordersList = append(ordersList, orderObj)
	}

	return ordersList, nil
}
