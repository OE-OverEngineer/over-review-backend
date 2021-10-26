package main

import (
	"github.com/OE-OverEngineer/over-review-backend/application/rest"
	"github.com/OE-OverEngineer/over-review-backend/configuration"
)

func main() {

	configuration, err := configuration.NewConfig()
	if err != nil {
		panic(err)
	}

	rest.Run(configuration)
}
