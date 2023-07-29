package handlerAuth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kbiits/gikslab_assesment/internal/app/constants"
	"github.com/kbiits/gikslab_assesment/internal/app/errors"
)

func (h *authHandler) Logout(ctx *fiber.Ctx) (err error) {
	token := ctx.Query(constants.TokenQuery)

	if token == "" {
		return errors.ErrUnauthorized
	}

	resp, err := h.authUc.Logout(ctx.Context(), token)
	if err != nil {
		return
	}

	return ctx.JSON(resp)
}
