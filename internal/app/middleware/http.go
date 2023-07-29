package middleware

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kbiits/gikslab_assesment/internal/app/config"
	"github.com/kbiits/gikslab_assesment/internal/app/constants"
	"github.com/kbiits/gikslab_assesment/internal/app/errors"
)

func RequestIdMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(constants.KeyRequestId, uuid.NewString())
		return c.Next()
	}
}

func BasicAuthMiddleware(cfg config.HttpConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authValue := c.Get("Authorization")
		if !strings.HasPrefix(authValue, "Basic ") {
			return errors.ErrUnauthorized
		}

		base64encodedCredentials := strings.TrimSpace(strings.ReplaceAll(authValue, "Basic ", ""))

		correctCredentials := fmt.Sprintf("%v:%v", cfg.BasicAuthUname, cfg.BasicAuthPassword)
		encodedCorrectCredentials := make([]byte, base64.StdEncoding.EncodedLen(len(correctCredentials)))
		base64.StdEncoding.Encode(encodedCorrectCredentials, []byte(correctCredentials))

		if subtle.ConstantTimeCompare(encodedCorrectCredentials, []byte(base64encodedCredentials)) == 0 {
			return errors.ErrUnauthorized
		}

		return c.Next()
	}
}
