package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/aggregate"
)

var (
	ErrCustomerNotFound       = errors.New("customer not found")
	ErrFailedToAddCustomer    = errors.New("failed to add customer")
	ErrFailedToUpdateCustomer = errors.New("failed to update customer")
)

type CustomerRepository interface {
	Find(ID uuid.UUID) (customer *aggregate.Customer, err error)
	FindAll() (customers []*aggregate.Customer, err error)
	Add(customer *aggregate.Customer) (err error)
	Update(customer *aggregate.Customer) (err error)
}
