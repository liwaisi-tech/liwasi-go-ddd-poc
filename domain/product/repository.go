package product

import (
	"github.com/google/uuid"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/aggregate"
)

type ProductRepository interface {
	Find(ID uuid.UUID) (product *aggregate.Product, err error)
	FindAll() (products []*aggregate.Product, err error)
	Store(product *aggregate.Product) (err error)
	Update(product *aggregate.Product) (err error)
	Delete(ID uuid.UUID) (err error)
}
