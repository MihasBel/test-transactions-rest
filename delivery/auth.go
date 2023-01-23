package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (r *REST) extractUserID(_ *fiber.Ctx) (uuid.UUID, bool) {
	return uuid.MustParse("c3bb416e-9a47-11ed-a8fc-0242ac120002"), true
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
