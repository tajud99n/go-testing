package locations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/tajud99n/go-testing/src/api/domain/location"
	"github.com/tajud99n/go-testing/src/api/utils/errors"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryId string) (*location.Country, *errors.APIError) {
	fmt.Println("Inside providers")
	response := rest.Get(fmt.Sprintf(urlGetCountry, countryId))
	if response == nil || response.Response == nil {
		return nil, &errors.APIError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient response when trying to get country %s", countryId),
		}
	}

	if response.StatusCode > 299 {
		var apiErr errors.APIError
		if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {
			return nil, &errors.APIError{
				Status: http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error response when getting country %s", countryId),
			}
		}
		return nil, &apiErr
	}

	var result location.Country
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, &errors.APIError{
			Status: http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryId),
		}
	}
	return &result, nil
}
