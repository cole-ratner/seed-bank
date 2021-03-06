package seed

import (
	"strings"
	"time"

	"seed-bank/entity"
)

//Service seed usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateSeed create a seed
func (s *Service) CreateSeed(category, name, vendor string, quantity int, unit string) (entity.ID, error) {
	b, err := entity.NewSeed(category, name, vendor, quantity, unit)
	if err != nil {
		return b.ID, err
	}
	return s.repo.Create(b)
}

//GetSeed get a seed
func (s *Service) GetSeed(id entity.ID) (*entity.Seed, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return b, nil
}

//SearchSeeds search seeds
func (s *Service) SearchSeeds(query string) ([]*entity.Seed, error) {
	seeds, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(seeds) == 0 {
		return nil, entity.ErrNotFound
	}
	return seeds, nil
}

//ListSeeds list seeds
func (s *Service) ListSeeds() ([]*entity.Seed, error) {
	seeds, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(seeds) == 0 {
		return nil, entity.ErrNotFound
	}
	return seeds, nil
}

//DeleteSeed Delete a seed
func (s *Service) DeleteSeed(id entity.ID) error {
	_, err := s.GetSeed(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//UpdateSeed Update a seed
func (s *Service) UpdateSeed(e *entity.Seed) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
