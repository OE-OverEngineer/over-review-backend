package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type Hash struct {
	secretKey string
}

type Hasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password string, hash string) bool
}

func NewHasher(secretKey string) Hasher {
	return &Hash{
		secretKey: secretKey,
	}
}

func (h *Hash) HashPassword(password string) (string, error) {
	salt := h.secretKey
	bytes, err := bcrypt.GenerateFromPassword([]byte(salt+password), 14)
	return string(bytes), err
}

func (h *Hash) CheckPasswordHash(password string, hash string) bool {
	salt := h.secretKey
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(salt+password))
	return err == nil
}
