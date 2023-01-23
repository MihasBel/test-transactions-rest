package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (r *REST) byID(ctx *fiber.Ctx) error {
	t, err := r.t.GetTransactionByID(ctx.Context(), uuid.UUID{})
	if err != nil {
		return fiber.NewError(404, "transaction not found")
	}
	return ctx.JSON(t)
}
