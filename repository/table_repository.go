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
	query := "SELECT id_table, table_number, isOccupied, location FROM tables;"
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

func (tr *TableRepository) CreateTable(table model.Table) (int, error) {
	var id int
	query, err := tr.conn.Prepare("INSERT INTO tables" +
		"(table_number, location)" +
		"VALUES ($1, $2) RETURNING id_table")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(table.Table_number, table.Location).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil

}
