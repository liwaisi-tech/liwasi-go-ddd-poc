package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/entity"
)

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists")
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name string, description string, price float64) (product *Product, err error) {
	item, err := entity.NewItem(name, description)
	if err != nil {
		return
	}
	product = &Product{
		item:     item,
		price:    price,
		quantity: 0,
	}
	return
}

// GetID returns the product ID
func (product *Product) GetID() (ID uuid.UUID) {
	return product.item.ID
}

// GetItem returns the product item
func (product *Product) GetItem() (item *entity.Item) {
	return product.item
}

// GetPrice returns the product price
func (product *Product) GetPrice() (price float64) {
	return product.price
}
