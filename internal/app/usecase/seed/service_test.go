package seed

import (
	"testing"
	"time"

	"seed-bank/entity"

	"github.com/stretchr/testify/assert"
)

func newTestSeed() *entity.Seed {
	return &entity.Seed{
		Category:      "Beans",
		Name:          "Kentucky Pole Bean",
		Vendor:        "Johnny's",
		Quantity:      15,
		UnitOfMeasure: "oz",
		CreatedAt:     time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u := newTestSeed()
	_, err := m.CreateSeed(u.Category, u.Name, u.Vendor, u.Quantity, u.UnitOfMeasure)
	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newTestSeed()
	u2 := newTestSeed()
	u2.Name = "Tavera Bush Bean"

	uID, _ := m.CreateSeed(u1.Category, u1.Name, u1.Vendor, u1.Quantity, u1.UnitOfMeasure)
	_, _ = m.CreateSeed(u2.Category, u2.Name, u2.Vendor, u2.Quantity, u2.UnitOfMeasure)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchSeeds("kentucky")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "Kentucky Pole Bean", c[0].Name)

		c, err = m.SearchSeeds("Some Bean")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.ListSeeds()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetSeed(uID)
		assert.Nil(t, err)
		assert.Equal(t, u1.Name, saved.Name)
	})
}

func Test_Update(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u := newTestSeed()
	id, err := m.CreateSeed(u.Category, u.Name, u.Vendor, u.Quantity, u.UnitOfMeasure)
	assert.Nil(t, err)
	saved, _ := m.GetSeed(id)
	saved.Name = "Kentucky Pole Bean"
	assert.Nil(t, m.UpdateSeed(saved))
	updated, err := m.GetSeed(id)
	assert.Nil(t, err)
	assert.Equal(t, "Kentucky Pole Bean", updated.Name)
}

func TestDelete(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newTestSeed()
	u2 := newTestSeed()
	u2ID, _ := m.CreateSeed(u2.Category, u2.Name, u2.Vendor, u2.Quantity, u2.UnitOfMeasure)

	err := m.DeleteSeed(u1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteSeed(u2ID)
	assert.Nil(t, err)
	_, err = m.GetSeed(u2ID)
	assert.Equal(t, entity.ErrNotFound, err)
}
