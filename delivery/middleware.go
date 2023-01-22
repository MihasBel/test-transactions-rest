package delivery

import (
	"github.com/MihasBel/test-transactions-rest/internal/app"
	"github.com/gofiber/fiber/v2"
)

func (r *REST) isAuth(c *fiber.Ctx) error {
	auth := c.GetReqHeaders()["Authorization"]
	if auth != app.Config.APIKey {
		return c.Next() //fiber.ErrUnauthorized
	}
	return c.Next()
}
