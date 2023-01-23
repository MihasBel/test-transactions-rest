package delivery

import (
	"net/http"

	"github.com/MihasBel/test-transactions-rest/model"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (r *REST) place(ctx *fiber.Ctx) error {
	uid, ok := r.extractUserID(ctx)
	if !ok {
		return fiber.NewError(401, "unauthorized")
	}
	a := model.Amount{}
	if err := ctx.BodyParser(&a); err != nil {
		log.Error().Err(err).Msg("error while decode body")
		return fiber.NewError(http.StatusBadRequest, "error while decode request body")
	}
	id, err := r.t.PlaceTransaction(ctx.Context(), a.Amount, uid)
	if err != nil {
		return fiber.NewError(404, err.Error())
	}
	return ctx.JSON(id)
}
