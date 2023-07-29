package handler

import (
	handlerActivity "github.com/kbiits/gikslab_assesment/internal/app/handler/activity"
	handlerAuth "github.com/kbiits/gikslab_assesment/internal/app/handler/auth"
	handlerSkill "github.com/kbiits/gikslab_assesment/internal/app/handler/skill"
	ucActivity "github.com/kbiits/gikslab_assesment/internal/app/usecase/activity"
	ucAuth "github.com/kbiits/gikslab_assesment/internal/app/usecase/auth"
	ucSkill "github.com/kbiits/gikslab_assesment/internal/app/usecase/skill"
)

type Handlers struct {
	AuthHandler     handlerAuth.AuthHandler
	SkillHandler    handlerSkill.SkillHandler
	ActivityHandler handlerActivity.ActivityHandler
}

func NewHandlers(
	authUc ucAuth.AuthUsecase,
	skillUc ucSkill.SkillUsecase,
	activityUc ucActivity.ActivityUsecase,
) *Handlers {
	return &Handlers{
		AuthHandler:     handlerAuth.NewAuthHandler(authUc),
		SkillHandler:    handlerSkill.NewSkillHandler(skillUc),
		ActivityHandler: handlerActivity.NewActivityHandler(activityUc),
	}
}
