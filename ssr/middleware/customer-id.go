package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
)

func CustomerIdMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Skip for public routes
		if c.Path() == "/public" {
			return next(c)
		}

		existingHeader, err := c.Cookie("X-Customer-Id")
		if err != nil {
			if err != http.ErrNoCookie {
				return err
			}
		}

		if existingHeader != nil {
			c.Set("customerId", existingHeader.Value)
			return next(c)
		}

		customerId := generateId()
		cookie := &http.Cookie{
			Name:  "X-Customer-Id",
			Value: customerId,
			Path:  "/",
		}

		c.SetCookie(cookie)
		c.Set("customerId", customerId)

		return next(c)
	}
}

func generateId() string {
	return random.String(32)
}
