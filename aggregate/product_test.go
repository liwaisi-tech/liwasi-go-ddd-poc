package aggregate_test

import (
	"testing"

	"github.com/liwasi-tech/liwasi-go-ddd-poc/aggregate"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/entity"
	"github.com/stretchr/testify/assert"
)

// TestNewProduct_InvalidItem_Name tests the NewProduct function when
// the item name is invalid.
func TestNewProduct_InvalidItem_Name(t *testing.T) {
	// Given
	name := ""
	description := "A test product"
	price := 1.99
	// When
	product, err := aggregate.NewProduct(name, description, price)
	// Then
	assert.Nil(t, product)
	assert.Equal(t, entity.ErrMissingValues, err)
}

// TestNewProduct_InvalidItem_Description tests the NewProduct function when
// the item description is invalid.
func TestNewProduct_InvalidItem_Description(t *testing.T) {
	// Given
	name := "Test Product"
	description := ""
	price := 1.99
	// When
	product, err := aggregate.NewProduct(name, description, price)
	// Then
	assert.Nil(t, product)
	assert.Equal(t, entity.ErrMissingValues, err)
}

// TestNewProduct_ValidItem tests the NewProduct function when
// the item is valid.
func TestNewProduct_ValidItem(t *testing.T) {
	// Given
	name := "Test Product"
	description := "A test product"
	price := 1.99
	// When
	product, err := aggregate.NewProduct(name, description, price)
	// Then
	assert.NotNil(t, product)
	assert.Nil(t, err)
	assert.Equal(t, name, product.GetItem().Name)
	assert.Equal(t, description, product.GetItem().Description)
	assert.Equal(t, price, product.GetPrice())
	assert.NotNil(t, product.GetID())
}
