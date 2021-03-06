package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

//User data
type User struct {
	ID        ID
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Seeds     []ID
}

//NewUser create a new user
func NewUser(email, password, firstName, lastName string) (*User, error) {
	u := &User{
		ID:        NewID(),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
	}
	pwd, err := generatePassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = pwd
	err = u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return u, nil
}

//AddSeed add a book
func (u *User) AddSeed(id ID) error {
	_, err := u.GetSeed(id)
	if err == nil {
		return ErrSeedAlreadyBorrowed
	}
	u.Seeds = append(u.Seeds, id)
	return nil
}

//RemoveSeed remove a book
func (u *User) RemoveSeed(id ID) error {
	for i, j := range u.Seeds {
		if j == id {
			u.Seeds = append(u.Seeds[:i], u.Seeds[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

//GetSeed get a book
func (u *User) GetSeed(id ID) (ID, error) {
	for _, v := range u.Seeds {
		if v == id {
			return id, nil
		}
	}
	return id, ErrNotFound
}

//Validate validate data
func (u *User) Validate() error {
	if u.Email == "" || u.FirstName == "" || u.LastName == "" || u.Password == "" {
		return ErrInvalidEntity
	}

	return nil
}

//ValidatePassword validate user password
func (u *User) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
