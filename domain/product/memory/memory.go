package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/aggregate"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]*aggregate.Product
	sync.Mutex
}

func NewMemoryProductRepository() (productRepository *MemoryProductRepository) {
	productRepository = &MemoryProductRepository{
		products: make(map[uuid.UUID]*aggregate.Product),
	}
	return
}

// Find returns a product by ID
func (productRepository *MemoryProductRepository) Find(ID uuid.UUID) (product *aggregate.Product, err error) {
	product, ok := productRepository.products[ID]
	if !ok {
		err = fmt.Errorf("product ID %s error: %w", ID, aggregate.ErrProductNotFound)
		return
	}
	return
}

// FindAll returns all products
func (productRepository *MemoryProductRepository) FindAll() (products []*aggregate.Product, err error) {
	for _, product := range productRepository.products {
		products = append(products, product)
	}
	return
}

// Store stores a product
func (productRepository *MemoryProductRepository) Store(product *aggregate.Product) (err error) {
	productRepository.Lock()
	defer productRepository.Unlock()
	// Validate if product already exists
	_, ok := productRepository.products[product.GetID()]
	if ok {
		err = fmt.Errorf("product ID %s error: %w", product.GetID(), aggregate.ErrProductAlreadyExists)
		return
	}
	productRepository.products[product.GetID()] = product
	return
}

// Update updates a product
func (productRepository *MemoryProductRepository) Update(product *aggregate.Product) (err error) {
	productRepository.Lock()
	defer productRepository.Unlock()
	// Validate if product exists
	_, ok := productRepository.products[product.GetID()]
	if !ok {
		err = fmt.Errorf("product ID %s error: %w", product.GetID(), aggregate.ErrProductNotFound)
		return
	}
	productRepository.products[product.GetID()] = product
	return
}

// Delete deletes a product by ID
func (productRepository *MemoryProductRepository) Delete(ID uuid.UUID) (err error) {
	productRepository.Lock()
	defer productRepository.Unlock()
	// Validate if product exists
	_, ok := productRepository.products[ID]
	if !ok {
		err = fmt.Errorf("product ID %s error: %w", ID, aggregate.ErrProductNotFound)
		return
	}
	delete(productRepository.products, ID)
	return
}
