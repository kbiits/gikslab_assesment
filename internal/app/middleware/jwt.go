package middleware

import (
	goerrors "errors"
	"log"

	"github.com/gofiber/fiber/v2"
	jwtLib "github.com/golang-jwt/jwt/v5"
	"github.com/kbiits/gikslab_assesment/internal/app/constants"
	"github.com/kbiits/gikslab_assesment/internal/app/errors"
	"github.com/kbiits/gikslab_assesment/internal/app/jwt"
)

func ValidateToken(jwtImpl *jwt.Jwt) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Query(constants.TokenQuery)
		if token == "" {
			return errors.ErrUnauthorized
		}

		isBlacklisted, err := jwtImpl.CheckIsTokenBlacklisted(c.UserContext(), token)
		if err != nil {
			return err
		}
		if isBlacklisted {
			return errors.ErrUnauthorized
		}

		tokenParsed, err := jwtLib.Parse(token, func(t *jwtLib.Token) (interface{}, error) {
			if !jwtImpl.IsValidMethod(t) {
				return nil, goerrors.New("unknown method")
			}

			return jwtImpl.GetSecretKey(), nil
		}, jwtLib.WithIssuedAt())
		if err != nil {
			log.Println(err)
			return errors.ErrUnauthorized
		}

		mapClaims, ok := tokenParsed.Claims.(jwtLib.MapClaims)
		if ok && tokenParsed.Valid {
			userId := mapClaims["userId"].(float64)
			exp := mapClaims["exp"].(float64)
			iat := mapClaims["iat"].(float64)
			c.Locals(constants.KeyLoggedInUser, jwt.Claims{
				UserId:    int(userId),
				Username:  mapClaims["username"].(string),
				Profile:   mapClaims["profile"].(string),
				ExpiredAt: int64(exp),
				IssuedAt:  int64(iat),
			})
		} else {
			return errors.ErrUnauthorized
		}

		return c.Next()
	}
}

func BoardUsersOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals(constants.KeyLoggedInUser).(jwt.Claims)
		if !ok {
			return errors.ErrUnauthorized
		}

		if !claims.IsProfileBoard() {
			return errors.ErrUnauthorized
		}

		return c.Next()
	}
}

func ExpertUsersOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals(constants.KeyLoggedInUser).(jwt.Claims)
		if !ok {
			return errors.ErrUnauthorized
		}

		if !claims.IsProfileExpert() {
			return errors.ErrUnauthorized
		}

		return c.Next()
	}
}
