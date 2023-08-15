package service

import (
	"errors"
	"fmt"
)

type UserService interface {
	UpdateUserStatus(userID string) error
}

type OrderService struct {
	userService UserService
}

type CreateOrderCmd struct {
	UserID     string
	ProductIDs []string
	PromoKey   string
}

func NewOrderService(us UserService) *OrderService {
	return &OrderService{
		userService: us,
	}
}

func (s *OrderService) CreateOrder(cmd CreateOrderCmd) error {
	if len(cmd.ProductIDs) == 0 {
		return errors.New("no products to create order")
	}

	// order creation

	err := s.userService.UpdateUserStatus(cmd.UserID)
	if err != nil {
		return fmt.Errorf("updating status: %w", err)
	}

	return nil
}
