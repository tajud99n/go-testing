package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/tajud99n/go-testing/src/api/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tajud99n/go-testing/src/api/domain/location"
	"github.com/tajud99n/go-testing/src/api/utils/errors"

	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	getCountryfunc func(countryId string) (*location.Country, *errors.APIError)
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

type locationServiceMock struct{}

func (*locationServiceMock) GetCountry(countryId string) (*location.Country, *errors.APIError) {
	return getCountryfunc(countryId)
}

func TestGetCountryNotFound(t *testing.T) {
	// Mock LocationService methods
	getCountryfunc = func(countryId string) (*location.Country, *errors.APIError) {
		return nil, &errors.APIError{Status: http.StatusNotFound, Message: "Country not found"}
	}
	services.LocationService = &locationServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "countryId", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, response.Code)

	var apiError errors.APIError
	err := json.Unmarshal(response.Body.Bytes(), &apiError)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiError.Status)
	assert.EqualValues(t, "Country not found", apiError.Message)
}

func TestGetCountryNoError(t *testing.T) {
	// Mock LocationService methods
	getCountryfunc = func(countryId string) (*location.Country, *errors.APIError) {
		return &location.Country{Id: "AR", Name: "Argentina"}, nil
	}
	services.LocationService = &locationServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "countryId", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusOK, response.Code)

	var country location.Country
	err := json.Unmarshal(response.Body.Bytes(), &country)
	assert.Nil(t, err)

	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}
