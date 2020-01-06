package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tajud99n/go-testing/src/api/utils/errors"

	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetCountryNotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found", "error": "not_found", "status": 404}`,
	})

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
