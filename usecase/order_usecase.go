package usecase

import (
	"api-blog-go/model"
	"api-blog-go/repository"
	"fmt"
)

type OrderUsecase struct {
	repository repository.OrderRepository
}

func NewOrderUsecase(repo repository.OrderRepository) OrderUsecase {
	return OrderUsecase{
		repository: repo,
	}
}

func (ou *OrderUsecase) CreateSection(id_table int) (int, error) {
	tableId, err := ou.repository.CreateSection(id_table)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return tableId, nil
}

func (ou *OrderUsecase) CreateOrder(productsList []model.ProductsRequestOrder, id_section string) (int, error) {
	// IMPLEMENTE
	orderId, err := ou.repository.CreateOrder(id_section)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	rowsEffected, err := ou.repository.CreateOrderItem(orderId, productsList)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return rowsEffected, nil
}
