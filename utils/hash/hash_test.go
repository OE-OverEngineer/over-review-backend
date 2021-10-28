package hash

import "testing"

func TestHashFunction(t *testing.T) {
	password := "123456"
	hasher := NewHasher("local")
	hashPassword, err := hasher.HashPassword(password)
	if err != nil {
		t.Fatal(`Error when hash password`)
	}
	if !hasher.CheckPasswordHash(password, hashPassword) {
		t.Fatal(`Hash function not match checkPasswordHash`)
	}
}
