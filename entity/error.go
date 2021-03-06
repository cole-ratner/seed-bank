package entity

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("Not found")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("Invalid entity")

//ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot Be Deleted")

//ErrNotEnoughSeeds cannot borrow
var ErrNotEnoughSeeds = errors.New("Not enough seeds")

//ErrSeedAlreadyBorrowed cannot borrow
var ErrSeedAlreadyBorrowed = errors.New("Seed already borrowed")

//ErrSeedNotBorrowed cannot return
var ErrSeedNotBorrowed = errors.New("Seed not borrowed")
