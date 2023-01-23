package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (r *REST) extractUserID(c *fiber.Ctx) (uuid.UUID, bool) {

	user, ok := c.Locals("user_id").(*jwt.Token)
	if !ok {
		return uuid.UUID{}, false
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, false
	}

	id, ok := claims["user_id"].(uuid.UUID)
	if !ok {
		return uuid.UUID{}, false
	}

	return id, true
}
