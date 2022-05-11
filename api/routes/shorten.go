package routes

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
	"time"
	"url-shortner/api/db"
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

	r2 := db.CreateClients(1)
	defer r2.Close()
	val, err := r2.Get(db.Context, c.RealIP()).Result()

	if err == redis.Nil {
		_ = r2.Set(db.Context, c.RealIP(), os.Getenv("API_QUOTA"), 30*time.Minute).Err()
	} else {
		val, _ = r2.Get(db.Context, c.RealIP()).Result()
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.TTL(db.Context, c.RealIP()).Result()

			return c.JSON(http.StatusServiceUnavailable, echo.Map{
				"error":           "rate limit exceeded",
				"limit_time_left": limit,
			})

		}
	}

	if !govalidator.IsURL(body.URL) {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "URL is not correct"})
	}

	if !helpers.RemoveDomainError(body.URL) {
		return c.JSON(http.StatusServiceUnavailable, echo.Map{"error": "error in finding domain"})
	}

	body.URL = helpers.EnforceHTTP(body.URL)

	r2.Decr(db.Context, c.RealIP())
}
