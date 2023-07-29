package handlerSkill

import "github.com/gofiber/fiber/v2"

func (h *skillHandler) ListSkill(ctx *fiber.Ctx) (err error) {
	resp, err := h.skillUc.ListSkill(ctx.Context())
	if err != nil {
		return
	}

	return ctx.JSON(resp)
}
