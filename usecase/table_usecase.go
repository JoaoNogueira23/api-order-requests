package usecase

import (
	"api-blog-go/model"
	"api-blog-go/repository"
	"fmt"
)

type TableUsecase struct {
	repository repository.TableRepository
}

func NewTableUsecase(repo repository.TableRepository) TableUsecase {
	return TableUsecase{
		repository: repo,
	}
}

func (tu *TableUsecase) GetTables() ([]model.Table, error) {
	return tu.repository.GetTables()
}

func (tu *TableUsecase) CreateTable(table model.Table) (model.Table, error) {

	tableId, err := tu.repository.CreateTable(table)
	if err != nil {
		fmt.Println(err)
		return model.Table{}, err
	}

	table.ID = tableId

	return table, nil

}

func (tu *TableUsecase) GetTableById(id_table int) (*model.Table, error) {
	table, err := tu.repository.GetTableById(id_table)

	if err != nil {
		return nil, err
	}

	return table, nil
}
