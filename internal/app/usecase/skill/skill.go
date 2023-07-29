package ucSkill

import (
	"context"

	dtoSkill "github.com/kbiits/gikslab_assesment/internal/app/dto/skill"
	repoSkill "github.com/kbiits/gikslab_assesment/internal/app/repository/skill"
)

type SkillUsecase interface {
	ListSkill(ctx context.Context) (dtoSkill.ListSkillDtoResponse, error)
}

func NewSkillUsecase(
	skillRepo *repoSkill.Queries,
) SkillUsecase {
	return &skillUsecase{
		skillRepo: skillRepo,
	}
}

type skillUsecase struct {
	skillRepo *repoSkill.Queries
}
