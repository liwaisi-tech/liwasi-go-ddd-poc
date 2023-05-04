package services_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/aggregate"
	customerMock "github.com/liwasi-tech/liwasi-go-ddd-poc/domain/customer/mock"
	productMock "github.com/liwasi-tech/liwasi-go-ddd-poc/domain/product/mock"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/services"
	"github.com/stretchr/testify/assert"
)

func init_test_products(t *testing.T) (products []*aggregate.Product) {

	catLitter, err := aggregate.NewProduct("Fresh Step", "Very good Cat Litter", 80.0)
	if err != nil {
		t.Fatalf("error creating product: %s", err)
	}
	food, err := aggregate.NewProduct("Agility Food", "Cat's Food from gods", 20.0)
	if err != nil {
		t.Fatalf("error creating product: %s", err)
	}
	catCookies, err := aggregate.NewProduct("Cat Cookies", "Cookies for cats", 10.0)
	if err != nil {
		t.Fatalf("error creating product: %s", err)
	}
	products = []*aggregate.Product{catLitter, food, catCookies}
	return
}

// TestOrderService_NewOrderService_Failed tests the NewOrderService function
// with a failed configuration
func TestOrderService_NewOrderService_Failed(t *testing.T) {
	//Mocking a failed order configuration function
	failedOrderConfiguration := func(orderService *services.OrderService) (err error) {
		return errors.New("failed configuration")
	}
	orderService, err := services.NewOrderService(
		failedOrderConfiguration,
	)
	assert.NotNil(t, err)
	assert.Nil(t, orderService)
}

// TestOrderService_NewOrderService_Success tests the NewOrderService function
// with a successful configuration
func TestOrderService_NewOrderService_Success(t *testing.T) {
	//Mocking a successful order configuration function
	successfulOrderConfiguration := func(orderService *services.OrderService) (err error) {
		return nil
	}
	orderService, err := services.NewOrderService(
		successfulOrderConfiguration,
	)
	assert.Nil(t, err)
	assert.NotNil(t, orderService)
}

// TestOrderService_NewOrderService_Success tests the NewOrderService function
// with a MemoryCustomerRepository configuration and a MemoryProductRepository
// configuration
func TestOrderService_NewOrderService_Success_MemoryRepositories(t *testing.T) {
	products := init_test_products(t)
	orderService, err := services.NewOrderService(
		services.WithMemoryCustomerRepository(),
		services.WithMemoryProductRepository(products),
	)
	assert.Nil(t, err)
	assert.NotNil(t, orderService)
}

// TestOrderService_WithProductRepository_Failed tests the WithProductRepository
// function with a failed product repository
func TestOrderService_WithProductRepository_Failed(t *testing.T) {
	//Mocking a failed product repository
	mockProductRepository := productMock.NewMockProductRepository()
	products := init_test_products(t)
	mockProductRepository.On("Store", products[0]).Return(errors.New("failed to store product"))
	orderService, err := services.NewOrderService(
		services.WithProductRepository(mockProductRepository, products),
	)
	assert.NotNil(t, err)
	assert.Nil(t, orderService)
}

// TestCreateOrder_Failed_Find_Customer tests the CreateOrder function with a
// failed customer find
func TestCreateOrder_Failed_Find_Customer(t *testing.T) {
	//Mocking a failed customer find
	mockCustomerRepository := customerMock.NewMockCustomerRepository()
	customerID := uuid.New()
	products := init_test_products(t)
	productsIDs := []uuid.UUID{products[0].GetID(), products[1].GetID(), products[2].GetID()}
	mockCustomerRepository.On("Find", customerID).Return(nil, errors.New("failed to find customer"))
	orderService, err := services.NewOrderService(
		services.WithCustomerRepository(mockCustomerRepository),
	)
	assert.Nil(t, err)
	assert.NotNil(t, orderService)
	err = orderService.CreateOrder(customerID, productsIDs)
	assert.NotNil(t, err)
}

// TestCreateOrder_Failed_Find_Product tests the CreateOrder function with a
// failed product find
func TestCreateOrder_Failed_Find_Product(t *testing.T) {
	//Mocking a failed product find
	mockProductRepository := productMock.NewMockProductRepository()
	mockCustomerRepository := customerMock.NewMockCustomerRepository()
	bruce, err := aggregate.NewCustomer("Bruce")
	assert.Nil(t, err)
	products := init_test_products(t)
	productsIDs := []uuid.UUID{products[0].GetID(), products[1].GetID(), products[2].GetID()}
	mockCustomerRepository.On("Find", bruce.GetID()).Return(bruce, nil)
	mockProductRepository.On("Store", products[0]).Return(nil)
	mockProductRepository.On("Store", products[1]).Return(nil)
	mockProductRepository.On("Store", products[2]).Return(nil)
	mockProductRepository.On("Find", products[0].GetID()).Return(nil, errors.New("failed to find product"))
	orderService, err := services.NewOrderService(
		services.WithCustomerRepository(mockCustomerRepository),
		services.WithProductRepository(mockProductRepository, products),
	)
	assert.Nil(t, err)
	assert.NotNil(t, orderService)
	err = orderService.CreateOrder(bruce.GetID(), productsIDs)
	assert.NotNil(t, err)
}

// TestCreateOrder_Success tests the CreateOrder function with a successful
// order creation
func TestCreateOrder_Success(t *testing.T) {
	mockProductRepository := productMock.NewMockProductRepository()
	mockCustomerRepository := customerMock.NewMockCustomerRepository()
	bruce, err := aggregate.NewCustomer("Bruce")
	assert.Nil(t, err)
	products := init_test_products(t)
	productsIDs := []uuid.UUID{products[0].GetID(), products[1].GetID(), products[2].GetID()}
	mockCustomerRepository.On("Find", bruce.GetID()).Return(bruce, nil)
	mockProductRepository.On("Store", products[0]).Return(nil)
	mockProductRepository.On("Store", products[1]).Return(nil)
	mockProductRepository.On("Store", products[2]).Return(nil)
	mockProductRepository.On("Find", products[0].GetID()).Return(products[0], nil)
	mockProductRepository.On("Find", products[1].GetID()).Return(products[1], nil)
	mockProductRepository.On("Find", products[2].GetID()).Return(products[2], nil)
	orderService, err := services.NewOrderService(
		services.WithCustomerRepository(mockCustomerRepository),
		services.WithProductRepository(mockProductRepository, products),
	)
	assert.Nil(t, err)
	assert.NotNil(t, orderService)
	err = orderService.CreateOrder(bruce.GetID(), productsIDs)
	assert.Nil(t, err)
}
