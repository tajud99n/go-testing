package app

import (
	"github.com/tajud99n/go-testing/src/api/controllers"
)

func mapURLs()  {
	router.GET("/locations/countries/:countryId", controllers.GetCountry)
}