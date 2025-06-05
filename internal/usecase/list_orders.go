package usecase

import (
	"github.com/Natayoane/go-clean-architecture/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (uc *ListOrdersUseCase) Execute() ([]*entity.Order, error) {
	orders, err := uc.OrderRepository.List()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
