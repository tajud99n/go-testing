package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tajud99n/go-testing/src/api/services"
)

func GetCountry(c *gin.Context) {
	country, err := services.GetCountry(c.Param("countryId"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, country)
}
