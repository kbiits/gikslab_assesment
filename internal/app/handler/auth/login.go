package handlerAuth

import (
	"github.com/gofiber/fiber/v2"
	dtoAuth "github.com/kbiits/gikslab_assesment/internal/app/dto/auth"
	errorAuth "github.com/kbiits/gikslab_assesment/internal/app/errors/auth"
)

func (h *authHandler) Login(ctx *fiber.Ctx) (err error) {
	req := new(dtoAuth.LoginDtoRequest)

	err = ctx.BodyParser(req)
	if err != nil {
		if err == fiber.ErrUnprocessableEntity {
			err = errorAuth.ErrLoginInvalid
		}

		return
	}

	err = validateLoginRequest(req)
	if err != nil {
		return
	}

	resp, err := h.authUc.Login(ctx.Context(), req)
	if err != nil {
		return
	}

	return ctx.JSON(resp)
}

func validateLoginRequest(req *dtoAuth.LoginDtoRequest) error {
	if req.Username == "" || req.Password == "" {
		return errorAuth.ErrLoginInvalid
	}

	return nil
}
