package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (r *REST) byID(ctx *fiber.Ctx) error {
	ids := ctx.Params("transactionId", "")
	id, err := uuid.Parse(ids)
	if err != nil {
		return fiber.NewError(404, "wrong id format")
	}
	t, err := r.t.GetTransactionByID(ctx.Context(), id)
	if err != nil {
		return fiber.NewError(404, "transaction not found")
	}
	return ctx.JSON(t)
}
