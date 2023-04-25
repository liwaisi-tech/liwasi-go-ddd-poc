package entity

import (
	"time"

	"github.com/google/uuid"
)

// Cat entity used to represent a cat
type Cat struct {
	ID        uuid.UUID
	Name      string
	BirthDate *time.Time
	Sterility bool
	Gender    string
}
