package main

import (
	"github.com/OE-OverEngineer/over-review-backend/utils/hash"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// configuration, err := configuration.NewConfig()
	// if err != nil {
	// 	panic(err)
	// }

	// rest.Run(configuration.Port)

	password := "test"
	hashPass, _ := hash.HashPassword(password)
	print(hashPass)
	print(hash.CheckPasswordHash(password, hashPass))
}
