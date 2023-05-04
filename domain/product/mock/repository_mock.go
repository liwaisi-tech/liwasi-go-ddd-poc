package mock

import (
	"github.com/google/uuid"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/aggregate"
	"github.com/stretchr/testify/mock"
)

type mockProductRepository struct {
	mock.Mock
}

// NewMockProductRepository creates a new mock for the ProductRepository interface
func NewMockProductRepository() *mockProductRepository {
	return &mockProductRepository{}
}

// product.ProductRepository interface mock implementation
func (m *mockProductRepository) Find(ID uuid.UUID) (product *aggregate.Product, err error) {
	args := m.Called(ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*aggregate.Product), args.Error(1)
}

func (m *mockProductRepository) FindAll() (products []*aggregate.Product, err error) {
	args := m.Called()
	return args.Get(0).([]*aggregate.Product), args.Error(1)
}

func (m *mockProductRepository) Store(product *aggregate.Product) (err error) {
	args := m.Called(product)
	return args.Error(0)
}

func (m *mockProductRepository) Update(product *aggregate.Product) (err error) {
	args := m.Called(product)
	return args.Error(0)
}

func (m *mockProductRepository) Delete(ID uuid.UUID) (err error) {
	args := m.Called(ID)
	return args.Error(0)
}
