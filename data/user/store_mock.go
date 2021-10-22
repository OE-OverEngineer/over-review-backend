package user

import (
	"errors"

	"github.com/OE-OverEngineer/over-review-backend/utils/hash"
)

type userRepositoryDB struct {
	users []User
}

func NewUserRepositoryMock() UserRepository {
	hasher := hash.NewHasher("local")
	hashPass, err := hasher.HashPassword("123456")
	if err != nil {
		panic(err)
	}
	users := []User{
		{ID: 1, FirstName: "Ponlawat", LastName: "Suparat", Email: "thekogo@gmail.com", Password: hashPass},
		{ID: 2, FirstName: "Nawakarn", LastName: "Leerattanachote", Email: "nawakarn@gmail.com", Password: hashPass},
		{ID: 3, FirstName: "Choolerk", LastName: "Taebanpakul", Email: "choolerk@gmail.com", Password: hashPass},
		{ID: 4, FirstName: "Kittapat", LastName: "Dechkul", Email: "kittapat@gmail.com", Password: hashPass},
	}
	return &userRepositoryDB{
		users: users,
	}
}

func (r *userRepositoryDB) GetAll() ([]User, error) {
	return r.users, nil
}

func (r *userRepositoryDB) GetUser(id int) (*User, error) {

	if id <= 0 {
		return nil, errors.New("id must be greater than 0")
	}

	for _, v := range r.users {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, nil
}
