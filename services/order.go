package services

import (
	"github.com/liwasi-tech/liwasi-go-ddd-poc/domain/customer"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/domain/customer/memory"
)

type OrderConfiguration func(orderService *OrderService) (err error)

type OrderService struct {
	customer customer.CustomerRepository
}

func NewOrderService(configurations ...OrderConfiguration) (orderService *OrderService, err error) {
	for _, config := range configurations {
		err = config(orderService)
		if err != nil {
			return
		}
	}
	return
}

// WithCustomerRepository sets the customer repository for the order service
func WithCustomerRepository(customerRepository customer.CustomerRepository) (configuration OrderConfiguration) {
	return func(orderService *OrderService) (err error) {
		orderService.customer = customerRepository
		return
	}
}

// WithMemoryCustomerRepository sets the customer repository for the order service
func WithMemoryCustomerRepository() (configuration OrderConfiguration) {
	repo := memory.NewMemoryCustomerRepository()
	return WithCustomerRepository(repo)
}

// CreateOrder creates an order for a customerID, with a list of products
