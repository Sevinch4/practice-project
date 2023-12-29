package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Price       int
	Category_id uuid.UUID
	Created_at  time.Time
	Updated_at  time.Time
}
