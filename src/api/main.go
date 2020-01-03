package main

import (
	"fmt"

	"github.com/tajud99n/go-testing/src/api/providers/locations"
)

func main() {
	country, err := locations.GetCountry("AR")

	fmt.Println(err)
	fmt.Println(country)
}
