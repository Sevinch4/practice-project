package models

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	ID         uuid.UUID
	Name       string
	Created_at time.Time
	Updated_at time.Time
}
