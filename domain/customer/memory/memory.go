package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/aggregate"
	repository "github.com/liwasi-tech/liwasi-go-ddd-poc/domain/customer"
)

type MemoryCustomerRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func NewMemoryCustomerRepository() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mcr *MemoryCustomerRepository) Find(ID uuid.UUID) (customer aggregate.Customer, err error) {
	customer, ok := mcr.customers[ID]
	if !ok {
		err = repository.ErrCustomerNotFound
		return
	}
	return
}

func (mcr *MemoryCustomerRepository) FindAll() (customers []aggregate.Customer, err error) {
	for _, customer := range mcr.customers {
		customers = append(customers, customer)
	}
	return
}

func (mcr *MemoryCustomerRepository) Add(customer aggregate.Customer) (err error) {
	if mcr.customers == nil {
		mcr.Lock()
		mcr.customers = make(map[uuid.UUID]aggregate.Customer)
		mcr.Unlock()
	}
	// Make sure the customer doesn't already exist
	if _, ok := mcr.customers[customer.GetID()]; ok {
		err = fmt.Errorf(
			"customer %s already exists: %w",
			customer.GetID(),
			repository.ErrFailedToAddCustomer,
		)
		return
	}
	// Add the customer
	mcr.Lock()
	mcr.customers[customer.GetID()] = customer
	mcr.Unlock()
	return
}

func (mcr *MemoryCustomerRepository) Update(customer aggregate.Customer) (err error) {
	if mcr.customers == nil {
		mcr.Lock()
		mcr.customers = make(map[uuid.UUID]aggregate.Customer)
		mcr.Unlock()
	}
	// Make sure the customer exists
	if _, ok := mcr.customers[customer.GetID()]; !ok {
		err = fmt.Errorf(
			"customer %s does not exist: %w",
			customer.GetID(),
			repository.ErrFailedToUpdateCustomer,
		)
		return
	}
	// Update the customer
	mcr.Lock()
	mcr.customers[customer.GetID()] = customer
	mcr.Unlock()
	return
}
