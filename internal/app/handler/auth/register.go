package handlerAuth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kbiits/gikslab_assesment/internal/app/constants"
	dtoAuth "github.com/kbiits/gikslab_assesment/internal/app/dto/auth"
	errorAuth "github.com/kbiits/gikslab_assesment/internal/app/errors/auth"
)

func (h *authHandler) Register(ctx *fiber.Ctx) (err error) {
	req := new(dtoAuth.RegisterDtoRequest)

	err = ctx.BodyParser(req)
	if err != nil {
		return
	}

	err = validateRegisterDto(req)
	if err != nil {
		return
	}

	resp, err := h.authUc.Register(ctx.Context(), req)
	if err != nil {
		return
	}

	return ctx.JSON(resp)
}

var mapAvailableProfile = map[string]bool{
	constants.ProfileBoard:      true,
	constants.ProfileExpert:     true,
	constants.ProfileTrainer:    true,
	constants.ProfileCompetitor: true,
}

func validateRegisterDto(req *dtoAuth.RegisterDtoRequest) error {

	if req.Email == "" || req.Name == "" || req.Password == "" || req.Username == "" || req.Profile == "" {
		return errorAuth.ErrRegisterFailed
	}

	if _, ok := mapAvailableProfile[req.Profile]; !ok {
		return errorAuth.ErrRegisterFailed
	}

	return nil
}
