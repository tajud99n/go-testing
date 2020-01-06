package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountriesNotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"status": 404, "message": "no country with id AR", "error": "not_found"}`,
	})
	response, err := http.Get("http://localhost:8080/locations/countries/AR")
	assert.Nil(t, err)
	assert.NotNil(t, response)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

	var apiErr errors.APIError
	err = json.Unmarshal(bytes, &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
}
