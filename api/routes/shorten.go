package routes

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"url-shortner/api/helpers"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"customShort"`
	ExpireTime  time.Duration `json:"expireTime"`
}

type response struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"customShort"`
	ExpireTime     time.Duration `json:"expireTime"`
	XRAteRemain    int           `json:"xRAteRemain"`
	XRestLimitRest time.Duration `json:"xRestLimitRest"`
}

func ShortenURL(c echo.Context) error {
	body := new(request)

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "cannot pars JSON"})
	}

	if !govalidator.IsURL(body.URL) {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "URL is not correct"})
	}

	if !helpers.RemoveDomainError(body.URL) {
		return c.JSON(http.StatusServiceUnavailable, echo.Map{"error": "error in finding domain"})
	}

	body.URL = helpers.EnforceHTTP(body.URL)

}

func ResolveURL(c echo.Context) error {

}
