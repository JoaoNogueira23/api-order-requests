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
