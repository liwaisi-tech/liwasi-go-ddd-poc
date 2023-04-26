package entity

import (
	"errors"

	"github.com/google/uuid"
)

// Item entity used to represent an item in the store. This item can be buyed by a user.
// In this case, the user is a cat.

var (
	ErrMissingValues = errors.New("missing values to create item")
)

type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}

// NewItem creates a new item
func NewItem(name string, description string) (item *Item, err error) {
	if name == "" || description == "" {
		err = ErrMissingValues
		return
	}
	item = &Item{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
	}
	return
}
