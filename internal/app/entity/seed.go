package entity

import (
	"time"
)

//Seed data
type Seed struct {
	ID            ID
	Category      string
	Name          string
	Vendor        string
	Quantity      int
	UnitOfMeasure string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

//NewSeed create a new seed
func NewSeed(category, name, vendor string, quantity int, unitOfMeasure string) (*Seed, error) {
	s := &Seed{
		ID:            NewID(),
		Name:          name,
		Category:      category,
		Vendor:        vendor,
		Quantity:      quantity,
		UnitOfMeasure: unitOfMeasure,
		CreatedAt:     time.Now(),
	}
	err := s.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return s, nil
}

//Validate validate seed
func (s *Seed) Validate() error {
	if s.Category == "" || s.Name == "" || s.Vendor == "" || s.Quantity <= 0 || s.UnitOfMeasure == "" {
		return ErrInvalidEntity
	}
	return nil
}
