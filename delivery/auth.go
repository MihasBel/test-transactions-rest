package delivery

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (r *REST) extractUserID(c *fiber.Ctx) (uuid.UUID, bool) {
	tokenString := strings.Split(c.GetReqHeaders()["Authorization"], "Bearer ")[1]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.cfg.Secret), nil
	})
	if err != nil {
		log.Error().Err(err).Msg("cannot get token")
		return uuid.UUID{}, false
	}
	ids, ok := claims["user_id"]
	if !ok {
		log.Error().Err(err).Msg("cannot find user_id in claims")
		return uuid.UUID{}, false
	}
	id, err := uuid.Parse(ids.(string))
	if err != nil {
		log.Error().Err(err).Msg("wrong format of user_id in claims")
		return uuid.UUID{}, false
	}
	return id, true
}
