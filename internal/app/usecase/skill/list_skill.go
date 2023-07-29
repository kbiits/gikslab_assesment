package ucSkill

import (
	"context"

	dtoSkill "github.com/kbiits/gikslab_assesment/internal/app/dto/skill"
)

func (uc *skillUsecase) ListSkill(ctx context.Context) (resp dtoSkill.ListSkillDtoResponse, err error) {
	resp = make(dtoSkill.ListSkillDtoResponse, 0)
	skills, err := uc.skillRepo.GetAllSkills(ctx)
	if err != nil {
		return
	}

	for _, skill := range skills {
		resp = append(resp, dtoSkill.SkillDtoResponse{
			Id:        skill.ID,
			SkillName: skill.Name,
		})
	}

	return
}
