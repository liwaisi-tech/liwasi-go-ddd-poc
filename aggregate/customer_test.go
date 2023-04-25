package aggregate_test

import (
	"testing"

	"github.com/liwasi-tech/liwasi-go-ddd-poc/aggregate"
	"github.com/stretchr/testify/assert"
)

// TestCustomer_NewCustomer_InvalidCat tests the NewCustomer function when
// the cat is invalid.
func TestCustomer_NewCustomer_InvalidCat(t *testing.T) {
	// Given
	name := ""
	// When
	customer, err := aggregate.NewCustomer(name)
	// Then
	assert.Nil(t, customer)
	assert.Equal(t, aggregate.ErrInvalidCat, err)
}

// TestCustomer_NewCustomer_ValidCat tests the NewCustomer function when
// the cat is valid.
func TestCustomer_NewCustomer_ValidCat(t *testing.T) {
	// Given
	name := "Phoney"
	// When
	customer, err := aggregate.NewCustomer(name)
	// Then
	assert.NotNil(t, customer)
	assert.Nil(t, err)
}
