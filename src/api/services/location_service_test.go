package services

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCountry(t *testing.T) {
	ls := locationService{}
	country, err := ls.GetCountry("AR")

	if country != nil {
		t.Error("Country should nil")
	}

	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to get country AR", err.Message)

	fmt.Println(err)
}
