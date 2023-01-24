package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (r *REST) extractUserID(c *fiber.Ctx) (uuid.UUID, bool) {
	return uuid.MustParse("bc708c12-7794-4716-b26c-47b4373d9716"), true //TODO processing jwt
	/*	user, ok := c.Locals("user_id").(*jwt.Token)
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

		return id, true*/
}
