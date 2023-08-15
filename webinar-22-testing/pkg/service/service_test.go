package service_test

import (
	"testing"
	"webinar-22/pkg/service"
)

type userServiceMock struct {
	f func(userID string) error
}

func (m userServiceMock) UpdateUserStatus(userID string) error {
	return m.f(userID)
}

func TestOrderServiceCreateOrder(t *testing.T) {
	inputCmd := service.CreateOrderCmd{
		UserID: "test-user-id",
		ProductIDs: []string{
			"test-product-id-1",
			"test-product-id-2",
		},
	}

	wantErr := false

	mock := userServiceMock{
		f: func(userID string) error {
			if userID != inputCmd.UserID {
				t.Logf("User service called with different user id: %v", userID)
				t.Fail()
			}

			return nil
		},
	}

	orderService := service.NewOrderService(mock)

	err := orderService.CreateOrder(inputCmd)
	if err != nil && !wantErr {
		t.Logf("Had unexpected error: %v", err)
		t.Fail()
	}
}
