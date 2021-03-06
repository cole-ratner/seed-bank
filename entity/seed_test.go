package entity_test

import (
	"testing"

	"seed-bank/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewSeed(t *testing.T) {
	s, err := entity.NewSeed("Corn", "Roy's Flint", "High Mowing", 100, "oz")
	assert.Nil(t, err)
	assert.Equal(t, s.Name, "Roy's Flint")
	assert.NotNil(t, s.ID)
}

func TestSeedValidate(t *testing.T) {
	type test struct {
		category string
		name     string
		vendor   string
		quantity int
		unit     string
		want     error
	}

	tests := []test{
		{
			category: "Corn",
			name:     "Roy's Flint",
			vendor:   "High Mowin",
			quantity: 100,
			unit:     "oz",
			want:     nil,
		},
		{
			category: "",
			name:     "Roy's Flint",
			vendor:   "High Mowin",
			quantity: 100,
			unit:     "oz",
			want:     entity.ErrInvalidEntity,
		},
		{
			category: "Corn",
			name:     "",
			vendor:   "High Mowin",
			quantity: 100,
			unit:     "oz",
			want:     entity.ErrInvalidEntity,
		},
		{
			category: "Corn",
			name:     "Roy's Flint",
			vendor:   "",
			quantity: 100,
			unit:     "oz",
			want:     entity.ErrInvalidEntity,
		},
		{
			category: "Corn",
			name:     "Roy's Flint",
			vendor:   "High Mowin",
			quantity: 0,
			unit:     "oz",
			want:     entity.ErrInvalidEntity,
		},
		{
			category: "Corn",
			name:     "Roy's Flint",
			vendor:   "High Mowin",
			quantity: 100,
			unit:     "",
			want:     entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {

		_, err := entity.NewSeed(tc.category, tc.name, tc.vendor, tc.quantity, tc.unit)
		assert.Equal(t, err, tc.want)
	}

}
