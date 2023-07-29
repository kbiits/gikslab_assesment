package handlerAuth

import (
	"github.com/gofiber/fiber/v2"
	authUc "github.com/kbiits/gikslab_assesment/internal/app/usecase/auth"
)

type authHandler struct {
	authUc authUc.AuthUsecase
}

type AuthHandler interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
}

func NewAuthHandler(authUc authUc.AuthUsecase) AuthHandler {
	return &authHandler{
		authUc: authUc,
	}
}
