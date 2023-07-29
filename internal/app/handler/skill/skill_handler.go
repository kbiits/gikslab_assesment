package handlerSkill

import (
	"github.com/gofiber/fiber/v2"
	ucSkill "github.com/kbiits/gikslab_assesment/internal/app/usecase/skill"
)

type skillHandler struct {
	skillUc ucSkill.SkillUsecase
}

type SkillHandler interface {
	ListSkill(ctx *fiber.Ctx) error
}

func NewSkillHandler(skillUc ucSkill.SkillUsecase) SkillHandler {
	return &skillHandler{
		skillUc: skillUc,
	}
}
