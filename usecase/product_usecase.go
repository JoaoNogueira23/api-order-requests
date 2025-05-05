package usecase

import (
	"api-blog-go/model"
	"api-blog-go/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts(page int, limit int) ([]model.Products, int, error) {
	return pu.repository.GetProducts(page, limit)
}

func (pu *ProductUsecase) CreateProduct(products []model.Products) (string, error) {

	_, err := pu.repository.CreateProduct(products)
	if err != nil {
		return "Erro ao criar o(s) produto(s)", err
	}

	return "produto(s) criados com sucesso!", nil
}

func (pu *ProductUsecase) GetProductById(id_product string) (*model.Products, error) {

	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
