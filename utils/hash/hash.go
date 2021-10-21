package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type Hash struct {
	secretKey string
}

type Hasher interface {
	HashPassword(string) (string, error)
	CheckPasswordHash(string, string) bool
}

func NewHasher(secretKey string) Hasher {
	return Hash{
		secretKey: secretKey,
	}
}

func (h Hash) TTT() {
	print("TETS")
}

func (h Hash) HashPassword(password string) (string, error) {
	salt := "test"
	bytes, err := bcrypt.GenerateFromPassword([]byte(salt+password), 14)
	return string(bytes), err
}

func (h Hash) CheckPasswordHash(password, hash string) bool {
	salt := "test"
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(salt+password))
	return err == nil
}
