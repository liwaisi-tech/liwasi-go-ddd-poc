package entity

import "github.com/google/uuid"

// Item entity used to represent an item in the store. This item can be buyed by a user.
// In this case, the user is a cat.
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       float64
}
