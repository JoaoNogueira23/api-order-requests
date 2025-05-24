package repository

import (
	"api-blog-go/model"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type TableRepository struct {
	conn *sql.DB
}

func NewTableRepository(conn *sql.DB) TableRepository {
	return TableRepository{
		conn: conn,
	}
}

func (pr *TableRepository) GetTables() ([]model.Table, error) {
	query := "SELECT id_table, table_number, status, location, capacity FROM tables;"
	rows, err := pr.conn.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Table{}, err
	}

	var tablesList []model.Table
	var tableObj model.Table

	for rows.Next() {
		err = rows.Scan(
			&tableObj.ID,
			&tableObj.Table_number,
			&tableObj.Status,
			&tableObj.Location,
			&tableObj.Capacity)
		if err != nil {
			fmt.Println(err)
			return []model.Table{}, err
		}

		tablesList = append(tablesList, tableObj)
	}

	rows.Close()

	return tablesList, nil
}

func (tr *TableRepository) CreateTable(table model.Table) (string, error) {
	var id string
	// variables
	// Define o seed para o gerador de números aleatórios
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Gera um ULID (ordenado lexicograficamente)
	id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	query, err := tr.conn.Prepare("INSERT INTO tables" +
		"(id_table, table_number, location, capacity)" +
		"VALUES ($1,$2, $3, $4) RETURNING id_table")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = query.QueryRow(id, table.Table_number, table.Location, table.Capacity).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	query.Close()
	return id, nil

}

func (tr *TableRepository) GetTableById(id_table string) (*model.Table, error) {
	query, err := tr.conn.Prepare("SELECT id_table, table_number, status, location, capacity FROM tables WHERE id_table = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var table model.Table

	err = query.QueryRow(id_table).Scan(
		&table.ID,
		&table.Table_number,
		&table.Status,
		&table.Location,
		&table.Capacity)

	if err != nil {
		if err == sql.ErrNoRows {
			// table nopt found
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &table, nil
}
