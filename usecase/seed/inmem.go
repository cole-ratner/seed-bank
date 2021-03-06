package seed

import (
	"strings"

	"seed-bank/entity"
)

//inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.Seed
}

//newInmem create new repository
func newInmem() *inmem {
	var m = map[entity.ID]*entity.Seed{}
	return &inmem{
		m: m,
	}
}

//Create a seed
func (r *inmem) Create(e *entity.Seed) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get a seed
func (r *inmem) Get(id entity.ID) (*entity.Seed, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update a seed
func (r *inmem) Update(e *entity.Seed) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search seeds
func (r *inmem) Search(query string) ([]*entity.Seed, error) {
	var d []*entity.Seed
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}
	return d, nil
}

//List seeds
func (r *inmem) List() ([]*entity.Seed, error) {
	var d []*entity.Seed
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a seed
func (r *inmem) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
