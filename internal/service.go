package internal

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Travellers struct {
	db Storage
}

func NewTravellers(db Storage) Travellers {
	return Travellers{db: db}
}

func (t Travellers) GetTraveller(ctx context.Context, id uuid.UUID) (Traveller, error) {
	res, err := t.db.GetTraveller(ctx, id)
	if err != nil {
		return Traveller{}, fmt.Errorf("failed to get traveller from db: %w", err)
	}

	return res, nil
}

func (t Travellers) CreateTraveller() {

}

func (t Travellers) DeleteTraveller() {

}
