package services

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M)  {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetCountryRestClientError(t *testing.T) {
	ls := locationService{}
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: 0,
	})
	country, err := ls.GetCountry("AR")

	if country != nil {
		t.Error("Country should nil")
	}

	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to get country AR", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	ls := locationService{}
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "AR", "name": "Argentina", "time_zone": "GMT-03:00"}`,

	})
	country, err := ls.GetCountry("AR")

	if err != nil {
		t.Error(err)
	}

	assert.EqualValues(t, "AR", country.Id)
}
