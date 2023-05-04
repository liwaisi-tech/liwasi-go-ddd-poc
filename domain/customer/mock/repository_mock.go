package mock

import (
	"github.com/google/uuid"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/aggregate"
	"github.com/stretchr/testify/mock"
)

type mockCustomerRepository struct {
	mock.Mock
}

// NewMockCustomerRepository creates a new mock customer repository
func NewMockCustomerRepository() *mockCustomerRepository {
	return &mockCustomerRepository{}
}

// customer.CustomerRepository interface mock implementation

func (m *mockCustomerRepository) Add(customer *aggregate.Customer) (err error) {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *mockCustomerRepository) Find(ID uuid.UUID) (customer *aggregate.Customer, err error) {
	args := m.Called(ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*aggregate.Customer), args.Error(1)
}

func (m *mockCustomerRepository) FindAll() (customers []*aggregate.Customer, err error) {
	args := m.Called()
	return args.Get(0).([]*aggregate.Customer), args.Error(1)
}

func (m *mockCustomerRepository) Delete(ID uuid.UUID) (err error) {
	args := m.Called(ID)
	return args.Error(0)
}

func (m *mockCustomerRepository) Update(customer *aggregate.Customer) (err error) {
	args := m.Called(customer)
	return args.Error(0)
}
