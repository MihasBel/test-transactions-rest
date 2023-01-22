package delivery

import (
	"net/http"
	"strconv"
	"time"

	"github.com/MihasBel/test-transactions-rest/model"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// Create godoc
// @Summary Post a new transaction from the received json contained user ID and balance delta
// @Accept json
// @Param request body model.PostParam true "Post parameters schema"
// @Success 200 {object} model.Transaction
// @Router / [post]
// @Security ApiKeyAuth
func (r REST) create(c *fiber.Ctx) error {
	p := model.PostParam{}
	if err := c.BodyParser(&p); err != nil {
		log.Error().Err(err).Msg("error while decode transaction")
		return fiber.NewError(http.StatusBadRequest, "error while decode request body")
	}
	t, err := r.t.Create(c.Context(), p.UserID, p.Delta)
	if err != nil {
		log.Error().Err(err).Msg("error while insert one transaction to db")
		return err
	}
	if err := c.JSON(t); err != nil {
		log.Error().Err(err).Msg("error while marshal inserted one transaction")
		return err
	}
	return nil
}

// Get godoc
// @Summary Retrieves balance with all transactions on given userID
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object} model.Balance
// @Router /history/{id} [get]
// @Security ApiKeyAuth
func (r *REST) history(c *fiber.Ctx) error {
	ids := c.Params("userId", "")
	b, err := r.t.History(c.Context(), ids)
	if err != nil {
		log.Error().Err(err).Msg("error while getting balance by id")
		return err
	}
	if err := c.JSON(b); err != nil {
		log.Error().Err(err).Msg("error while marshal balance")
		return err
	}
	return nil
}

// Get godoc
// @Summary Retrieve balance only on given userID
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object} model.Balance
// @Router /balance/{id} [get]
// @Security ApiKeyAuth
func (r *REST) balance(c *fiber.Ctx) error {
	ids := c.Params("userId", "")
	b, err := r.t.Balance(c.Context(), ids)
	if err != nil {
		log.Error().Err(err).Msg("error while getting balance by id")
		return err
	}
	if err := c.JSON(b); err != nil {
		log.Error().Err(err).Msg("error while marshal balance")
		return err
	}
	return nil
}

// Get godoc
// @Summary Retrieves transaction on given ID
// @Produce json
// @Param id path string true "transaction ID"
// @Success 200 {object} model.Transaction
// @Router /byID/{id} [get]
// @Security ApiKeyAuth
func (r *REST) byID(c *fiber.Ctx) error {
	ids := c.Params("transactionId", "")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid ID supplied")
	}
	t, err := r.t.ByID(c.Context(), id)
	if err != nil {
		log.Error().Err(err).Msg("error while getting transaction by id")
		return err
	}
	if err := c.JSON(t); err != nil {
		log.Error().Err(err).Msg("error while marshal transaction")
		return err
	}
	return nil
}

// Get godoc
// @Summary Retrieves transaction on given user ID and time
// @Produce json
// @Param userId query string true "userId"
// @Param tTime query string true "tTime"
// @Success 200 {object} model.Transaction
// @Router /byTime/ [get]
// @Security ApiKeyAuth
func (r *REST) byTime(c *fiber.Ctx) error {
	ids := c.Params("userId", "")
	timeS := c.Params("tTime", "")
	tTime, err := time.Parse(timeS, timeS)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid time supplied")
	}
	t, err := r.t.ByTime(c.Context(), ids, tTime)
	if err != nil {
		log.Error().Err(err).Msg("error while getting transaction by time")
		return err
	}
	if err := c.JSON(t); err != nil {
		log.Error().Err(err).Msg("error while marshal transaction")
		return err
	}
	return nil
}
