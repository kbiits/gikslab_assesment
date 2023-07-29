package ucAuth

import (
	"context"
	"log"

	dtoAuth "github.com/kbiits/gikslab_assesment/internal/app/dto/auth"
	errorAuth "github.com/kbiits/gikslab_assesment/internal/app/errors/auth"
	repoSkill "github.com/kbiits/gikslab_assesment/internal/app/repository/skill"
	repoUser "github.com/kbiits/gikslab_assesment/internal/app/repository/user"
	"github.com/kbiits/gikslab_assesment/internal/app/utils"
)

func (uc *authUsecase) Register(ctx context.Context, req *dtoAuth.RegisterDtoRequest) (resp dtoAuth.RegisterDtoResponse, err error) {

	var skill *repoSkill.Skill
	if req.Skill != "" {
		skillModel, err := uc.skillRepo.FindSkillByName(ctx, req.Skill)
		if err != nil {
			return resp, errorAuth.ErrRegisterFailed
		}

		skill = &skillModel
	}

	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Println(err)
		return resp, errorAuth.ErrRegisterFailed
	}

	registerParams := repoUser.RegisterUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Username: req.Username,
		Password: hashedPass,
		Profile:  req.Profile,
	}

	if skill != nil {
		err = registerParams.SkillID.Scan(skill.ID)
		if err != nil {
			log.Println(err)
			return resp, errorAuth.ErrRegisterFailed
		}
	}

	_, err = uc.userRepo.RegisterUser(ctx, registerParams)

	if err != nil {
		log.Println(err)
		return resp, errorAuth.ErrRegisterFailed
	}

	resp = dtoAuth.NewRegisterDtoResponseSuccess()
	return
}
