package internal

import (
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	GetTraveller(ctx context.Context, id uuid.UUID) (Traveller, error)
}

type Traveller struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Age       int
}
