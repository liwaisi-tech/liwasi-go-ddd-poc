package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/entity"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/valueobject"
)

var (
	ErrInvalidCat = errors.New("invalid cat")
)

type Customer struct {
	// Cat is the lead entity in this aggregate
	// It means cat.ID is the main identifier for the customer
	cat          *entity.Cat
	products     []*entity.Item
	transactions []valueobject.Transaction
}

func NewCustomer(name string) (customer *Customer, err error) {
	if name == "" {
		err = ErrInvalidCat
		return
	}
	cat := &entity.Cat{
		ID:   uuid.New(),
		Name: name,
	}
	customer = &Customer{
		cat:          cat,
		products:     []*entity.Item{},
		transactions: []valueobject.Transaction{},
	}
	return
}

func (c *Customer) GetID() uuid.UUID {
	return c.cat.ID
}

func (c *Customer) SetID(ID uuid.UUID) {
	if c.cat == nil {
		c.cat = &entity.Cat{
			ID: ID,
		}
		return
	}
	c.cat.ID = ID
}

func (c *Customer) GetName() string {
	return c.cat.Name
}

func (c *Customer) SetName(name string) {
	c.cat.Name = name
}
