package usecase

import (
	"api-blog-go/model"
	"api-blog-go/repository"
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
