package memory_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/aggregate"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/domain/product/memory"
	"github.com/stretchr/testify/assert"
)

// TestMemoryProductRepository_Find_Failed tests the Find method of the MemoryProductRepository
// when the product cannot be found.
func TestMemoryProductRepository_Find_Failed(t *testing.T) {
	// Create a new product
	product, err := aggregate.NewProduct("Peporris", "It's the main food of the cat's god", 1.99)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	// Create a new repository
	repo := memory.NewMemoryProductRepository()
	assert.NotNil(t, repo)
	// Find the product
	foundProduct, err := repo.Find(product.GetID())
	assert.NotNil(t, err)
	assert.Nil(t, foundProduct)
}

// TestMemoryProductRepository_Find_Success tests the Find method of the MemoryProductRepository
// when the product can be found.
func TestMemoryProductRepository_Find_Success(t *testing.T) {
	// Create a new product
	product, err := aggregate.NewProduct("Peporris", "It's the main food of the cat's god", 1.99)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	// Create a new repository
	repo := memory.NewMemoryProductRepository()
	assert.NotNil(t, repo)
	// Store the product
	err = repo.Store(product)
	assert.Nil(t, err)
	// Find the product
	foundProduct, err := repo.Find(product.GetID())
	assert.Nil(t, err)
	assert.NotNil(t, foundProduct)
	assert.Equal(t, product, foundProduct)
}

// TestMemoryProductRepository_FindAll_Nil tests the FindAll method of the MemoryProductRepository
// when there are no products.
func TestMemoryProductRepository_FindAll_Nil(t *testing.T) {
	// Create a new repository
	repo := memory.NewMemoryProductRepository()
	assert.NotNil(t, repo)
	// Find all products
	products, err := repo.FindAll()
	assert.Nil(t, err)
	assert.Nil(t, products)
}

// TestMemoryProductRepository_FindAll_Success tests the FindAll method of the MemoryProductRepository
// when there are products.
func TestMemoryProductRepository_FindAll_Success(t *testing.T) {
	// Create a new product
	product, err := aggregate.NewProduct("Peporris", "It's the main food of the cat's god", 1.99)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	// Create a new repository
	repo := memory.NewMemoryProductRepository()
	assert.NotNil(t, repo)
	// Store the product
	err = repo.Store(product)
	assert.Nil(t, err)
	// Create a new product
	product, err = aggregate.NewProduct("Latorris", "It's the special food of the cat's god", 1.99)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	// Store the product
	err = repo.Store(product)
	assert.Nil(t, err)
	// Find all products
	products, err := repo.FindAll()
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 2, len(products))
}

// TestMemoryProductRepository_Store_Failed tests the Store method of the MemoryProductRepository
// when the product already exists.
func TestMemoryProductRepository_Store_Failed(t *testing.T) {
	// Create a new product
	product, err := aggregate.NewProduct("Peporris", "It's the main food of the cat's god", 1.99)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	// Create a new repository
	repo := memory.NewMemoryProductRepository()
	assert.NotNil(t, repo)
	// Store the product
	err = repo.Store(product)
	assert.Nil(t, err)
	// Store the product again
	err = repo.Store(product)
	assert.NotNil(t, err)
}

// TestMemoryProductRepository_Store_Success tests the Store method of the MemoryProductRepository
// when the product does not exist.
func TestMemoryProductRepository_Store_Success(t *testing.T) {
	// Create a new product
	product, err := aggregate.NewProduct("Peporris", "It's the main food of the cat's god", 1.99)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	// Create a new repository
	repo := memory.NewMemoryProductRepository()
	assert.NotNil(t, repo)
	// Store the product
	err = repo.Store(product)
	assert.Nil(t, err)
}

// TestMemoryProductRepository_Update_Failed tests the Update method of the MemoryProductRepository
// when the product does not exist.
func TestMemoryProductRepository_Update_Failed(t *testing.T) {
	// Create a new product
	product, err := aggregate.NewProduct("Peporris", "It's the main food of the cat's god", 1.99)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	// Create a new repository
	repo := memory.NewMemoryProductRepository()
	assert.NotNil(t, repo)
	// Update the product
	err = repo.Update(product)
	assert.NotNil(t, err)
}

// TestMemoryProductRepository_Update_Success tests the Update method of the MemoryProductRepository
// when the product exists.
func TestMemoryProductRepository_Update_Success(t *testing.T) {
	// Create a new product
	product, err := aggregate.NewProduct("Peporris", "It's the main food of the cat's god", 1.99)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	// Create a new repository
	repo := memory.NewMemoryProductRepository()
	assert.NotNil(t, repo)
	// Store the product
	err = repo.Store(product)
	assert.Nil(t, err)
	// Update the product
	err = repo.Update(product)
	assert.Nil(t, err)
}

// TestMemoryProductRepository_Delete_Failed tests the Delete method of the MemoryProductRepository
// when the product does not exist.
func TestMemoryProductRepository_Delete_Failed(t *testing.T) {
	// Create a new repository
	repo := memory.NewMemoryProductRepository()
	assert.NotNil(t, repo)
	// Delete the product
	err := repo.Delete(uuid.New())
	assert.NotNil(t, err)
}

// TestMemoryProductRepository_Delete_Success tests the Delete method of the MemoryProductRepository
// when the product exists.
func TestMemoryProductRepository_Delete_Success(t *testing.T) {
	// Create a new product
	product, err := aggregate.NewProduct("Peporris", "It's the main food of the cat's god", 1.99)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	// Create a new repository
	repo := memory.NewMemoryProductRepository()
	assert.NotNil(t, repo)
	// Store the product
	err = repo.Store(product)
	assert.Nil(t, err)
	// Delete the product
	err = repo.Delete(product.GetID())
	assert.Nil(t, err)
	// Find the product
	_, err = repo.Find(product.GetID())
	assert.NotNil(t, err)
}
