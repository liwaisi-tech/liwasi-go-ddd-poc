package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction entity is a value object used to represent a transaction in the store.
type Transaction struct {
	id        uuid.UUID
	from      uuid.UUID
	to        uuid.UUID
	amount    float64
	createdAt time.Time
}
