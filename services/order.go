package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/aggregate"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/domain/customer"
	customerMemory "github.com/liwasi-tech/liwasi-go-ddd-poc/domain/customer/memory"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/domain/product"
	productMemory "github.com/liwasi-tech/liwasi-go-ddd-poc/domain/product/memory"
)

type OrderConfiguration func(orderService *OrderService) (err error)

type OrderService struct {
	customer customer.CustomerRepository
	product  product.ProductRepository
}

// NewOrderService creates a new order service
func NewOrderService(cfgs ...OrderConfiguration) (orderService *OrderService, err error) {
	orderService = &OrderService{}
	for _, cfg := range cfgs {
		err = cfg(orderService)
		if err != nil {
			orderService = nil
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
	repo := customerMemory.NewMemoryCustomerRepository()
	return WithCustomerRepository(repo)
}

// WithProductRepository sets the product repository for the order service
func WithProductRepository(
	productRepository product.ProductRepository,
	products []*aggregate.Product,
) (configuration OrderConfiguration) {
	return func(orderService *OrderService) (err error) {
		for _, product := range products {
			err = productRepository.Store(product)
			if err != nil {
				return
			}
		}
		orderService.product = productRepository
		return
	}
}

// WithMemoryProductRepository sets the product repository for the order service
func WithMemoryProductRepository(products []*aggregate.Product) (configuration OrderConfiguration) {
	return func(orderService *OrderService) (err error) {
		productRepository := productMemory.NewMemoryProductRepository()
		return WithProductRepository(productRepository, products)(orderService)
	}
}

// CreateOrder creates a new order for a customer
func (orderService *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (err error) {
	// Get customer
	customerFound, err := orderService.customer.Find(customerID)
	if err != nil {
		return
	}
	var customerProducts []*aggregate.Product
	var total float64
	// Get products
	for _, productID := range productsIDs {
		product, errFind := orderService.product.Find(productID)
		if errFind != nil {
			err = errFind
			return
		}
		customerProducts = append(customerProducts, product)
		total += product.GetPrice()
	}
	log.Printf("Customer %s has orderd %d products", customerFound.GetName(), len(customerProducts))
	return
}
