package memory_test

import (
	"testing"

	"github.com/liwasi-tech/liwasi-go-ddd-poc/aggregate"
	"github.com/liwasi-tech/liwasi-go-ddd-poc/domain/customer/memory"
	"github.com/stretchr/testify/assert"
)

// TestCustomerMemoryRepository_Add_OK tests the Add method of the MemoryCustomerRepository
// when everything is OK.
func TestCustomerMemoryRepository_Add_OK(t *testing.T) {
	// Create a new customer
	customer, err := aggregate.NewCustomer("Bruch")
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	// Create a new repository
	repo := memory.NewMemoryCustomerRepository()
	assert.NotNil(t, repo)
	// Add the customer first time
	err = repo.Add(*customer)
	assert.Nil(t, err)
}

// TestCustomerMemoryRepository_Add_Failed tests the Add method of the MemoryCustomerRepository
// when the customer cannot be added.
func TestCustomerMemoryRepository_Add_Failed(t *testing.T) {
	// Create a new customer
	customer, err := aggregate.NewCustomer("Bruch")
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	// Create a new repository
	repo := memory.NewMemoryCustomerRepository()
	assert.NotNil(t, repo)
	// Add the customer first time
	err = repo.Add(*customer)
	assert.Nil(t, err)
	// Add the customer second time
	err = repo.Add(*customer)
	assert.NotNil(t, err)
}

// TestCustomerMemoryRepository_Find_OK tests the Find method of the MemoryCustomerRepository
// when everything is OK.
func TestCustomerMemoryRepository_Find_OK(t *testing.T) {
	// Create a new customer
	customer, err := aggregate.NewCustomer("Bruch")
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	// Create a new repository
	repo := memory.NewMemoryCustomerRepository()
	assert.NotNil(t, repo)
	// Add the customer first time
	err = repo.Add(*customer)
	assert.Nil(t, err)
	// Find the customer
	foundCustomer, err := repo.Find(customer.GetID())
	assert.Nil(t, err)
	assert.NotNil(t, foundCustomer)
}

// TestCustomerMemoryRepository_Find_Failed tests the Find method of the MemoryCustomerRepository
// when the customer cannot be found.
func TestCustomerMemoryRepository_Find_Failed(t *testing.T) {
	// Create a new customer
	customer, err := aggregate.NewCustomer("Bruch")
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	// Create a new repository
	repo := memory.NewMemoryCustomerRepository()
	assert.NotNil(t, repo)
	// Find the customer
	foundCustomer, err := repo.Find(customer.GetID())
	assert.NotNil(t, err)
	assert.Equal(t, aggregate.Customer{}, foundCustomer)
}

// TestCustomerMemoryRepository_Find_Not_Initialized tests the Find method of the MemoryCustomerRepository
// when the map is not initialized.
func TestCustomerMemoryRepository_Find_Not_Initialized(t *testing.T) {
	// Create a new repository
	repo := memory.MemoryCustomerRepository{}
	// Create a new customer
	customer, err := aggregate.NewCustomer("Bruch")
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	// Add the customer first time
	err = repo.Add(*customer)
	assert.Nil(t, err)
}

// TestCustomerMemoryRepository_Update_OK tests the Update method of the MemoryCustomerRepository
// when everything is OK.
func TestCustomerMemoryRepository_Update_OK(t *testing.T) {
	// Create a new customer
	customer, err := aggregate.NewCustomer("Bruch")
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	// Create a new repository
	repo := memory.NewMemoryCustomerRepository()
	assert.NotNil(t, repo)
	// Add the customer first time
	err = repo.Add(*customer)
	assert.Nil(t, err)
	// Update the customer
	customer.SetName("Ulce")
	err = repo.Update(*customer)
	assert.Nil(t, err)
	found, err := repo.Find(customer.GetID())
	assert.Nil(t, err)
	assert.Equal(t, "Ulce", found.GetName())
}

// TestCustomerMemoryRepository_Update_Failed tests the Update method of the MemoryCustomerRepository
// when everything is OK.
func TestCustomerMemoryRepository_Update_Failed(t *testing.T) {
	// Create a new customer
	customer, err := aggregate.NewCustomer("Bruch")
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	// Create a new repository
	repo := memory.MemoryCustomerRepository{}
	// Update the non existing customer
	newCustomer, err := aggregate.NewCustomer("Ulce")
	assert.Nil(t, err)
	err = repo.Update(*newCustomer)
	assert.NotNil(t, err)
}

// TestCustomerMemoryRepository_FindAll_OK_Empty tests the FindAll method of the MemoryCustomerRepository
// when everything is OK and the repository is empty.
func TestCustomerMemoryRepository_FindAll_OK_Empty(t *testing.T) {
	// Create a new repository
	repo := memory.NewMemoryCustomerRepository()
	assert.NotNil(t, repo)
	// Find all customers
	customers, err := repo.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(customers))
}

// TestCustomerMemoryRepository_FindAll_OK tests the FindAll method of the MemoryCustomerRepository
// when everything is OK.
func TestCustomerMemoryRepository_FindAll_OK(t *testing.T) {
	// Create a new repository
	repo := memory.NewMemoryCustomerRepository()
	assert.NotNil(t, repo)
	// Create a new customer
	customer, err := aggregate.NewCustomer("Bruch")
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	// Add the customer
	err = repo.Add(*customer)
	assert.Nil(t, err)
	// Create a new customer
	customer, err = aggregate.NewCustomer("Ulce")
	assert.Nil(t, err)
	// Add the customer
	err = repo.Add(*customer)
	assert.Nil(t, err)
	// Find all customers
	customers, err := repo.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(customers))
}
