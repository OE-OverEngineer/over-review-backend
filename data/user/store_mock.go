package user

import (
	"time"

	"github.com/OE-OverEngineer/over-review-backend/utils/hash"
	"gorm.io/gorm"
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
		{
			Model: gorm.Model{
				ID: 1,
			},
			FirstName:   "Ponlawat",
			LastName:    "Suparat",
			DisplayName: "",
			Avatar:      "",
			BirthDate:   time.Time{},
			Email:       "thekogo@gmail.com",
			Password:    hashPass,
			// Gender:      "Male",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			FirstName:   "Nawakarn",
			LastName:    "Leerattanachote",
			DisplayName: "",
			Avatar:      "",
			BirthDate:   time.Time{},
			Email:       "nawakarn@gmail.com",
			Password:    hashPass,
			// Gender:      "Female",
		},
		{
			Model: gorm.Model{
				ID: 3,
			},
			FirstName:   "Choolerk",
			LastName:    "Taebanpakul",
			DisplayName: "",
			Avatar:      "",
			BirthDate:   time.Time{},
			Email:       "guide@gmail.com",
			Password:    hashPass,
			// Gender:      "Male",
		},
		{
			Model: gorm.Model{
				ID: 4,
			},
			FirstName:   "Kittapat",
			LastName:    "Dechkul",
			DisplayName: "",
			Avatar:      "",
			BirthDate:   time.Time{},
			Email:       "kittapat@gmail.com",
			Password:    hashPass,
			// Gender:      "Male",
		},
	}
	return &userRepositoryDB{
		users: users,
	}
}

func (r *userRepositoryDB) GetAll() ([]User, error) {
	return r.users, nil
}

func (r *userRepositoryDB) GetUser(id int) (*User, error) {

	for _, v := range r.users {
		if int(v.Model.ID) == id {
			return &v, nil
		}
	}

	return nil, nil
}

func (r *userRepositoryDB) GetByToken() (*User, error) {

	// for _, v := range r.users {
	// 	if int(v.Model.ID) == id {
	// 		return &v, nil
	// 	}
	// }

	return nil, nil
}
func (r *userRepositoryDB) DeleteUser(id int) error {
	return nil
}
