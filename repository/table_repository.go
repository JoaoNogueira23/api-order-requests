package repository

import (
	"api-blog-go/model"
	"database/sql"
	"fmt"
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
	query := "SELECT id_table, table_number, isoccupied, location FROM tables;"
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
			&tableObj.IsOccupied,
			&tableObj.Location)
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
	query, err := tr.conn.Prepare("INSERT INTO tables" +
		"(table_number, location)" +
		"VALUES ($1, $2) RETURNING id_table")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = query.QueryRow(table.Table_number, table.Location).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	query.Close()
	return id, nil

}

func (tr *TableRepository) GetTableById(id_table int) (*model.Table, error) {
	query, err := tr.conn.Prepare("SELECT * FROM tables WHERE id_table = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var table model.Table

	err = query.QueryRow(id_table).Scan(
		&table.ID,
		&table.Table_number,
		&table.IsOccupied,
		&table.Location)

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
