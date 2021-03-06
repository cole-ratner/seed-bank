package seed

import (
	"seed-bank/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Seed, error)
	Search(query string) ([]*entity.Seed, error)
	List() ([]*entity.Seed, error)
}

//Writer Seed writer
type Writer interface {
	Create(e *entity.Seed) (entity.ID, error)
	Update(e *entity.Seed) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetSeed(id entity.ID) (*entity.Seed, error)
	SearchSeeds(query string) ([]*entity.Seed, error)
	ListSeeds() ([]*entity.Seed, error)
	CreateSeed(title string, author string, pages int, quantity int) (entity.ID, error)
	UpdateSeed(e *entity.Seed) error
	DeleteSeed(id entity.ID) error
}
