package services

import (
	"fmt"

	"github.com/tajud99n/go-testing/src/api/domain/location"
	"github.com/tajud99n/go-testing/src/api/providers/locations"
	"github.com/tajud99n/go-testing/src/api/utils/errors"
)

type locationService struct{}

type locationServiceInterface interface {
	GetCountry(countryId string) (*location.Country, *errors.APIError)
}

var (
	LocationService locationServiceInterface
)

func init() {
	LocationService = &locationService{}
}

func (s *locationService) GetCountry(countryId string) (*location.Country, *errors.APIError) {
	fmt.Println("Inside service")
	return locations.GetCountry(countryId)
}
