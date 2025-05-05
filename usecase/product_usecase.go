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

func (pu *ProductUsecase) CreateProduct(product model.Products) (model.Products, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Products{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(id_product string) (*model.Products, error) {

	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
