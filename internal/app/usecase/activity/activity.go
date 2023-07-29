package ucActivity

import (
	"context"

	dtoActivity "github.com/kbiits/gikslab_assesment/internal/app/dto/activity"
	repoActivity "github.com/kbiits/gikslab_assesment/internal/app/repository/activity"
	repoSkill "github.com/kbiits/gikslab_assesment/internal/app/repository/skill"
	repoUser "github.com/kbiits/gikslab_assesment/internal/app/repository/user"
)

type ActivityUsecase interface {
	AddActivity(ctx context.Context, req *dtoActivity.AddActivityDtoRequest) (dtoActivity.AddActivityDtoResponse, error)
	UpdateActivity(ctx context.Context, req *dtoActivity.AddActivityDtoRequest, actId int) (dtoActivity.UpdateActivityDtoResponse, error)
	DeleteActivity(ctx context.Context, actId int) (dtoActivity.DeleteActivityDtoResponse, error)
	ListActivityBySkillId(ctx context.Context, req *dtoActivity.ListActivitiesDtoRequest) (resp dtoActivity.ListActivitiesDtoResponse, err error)
}

func NewActivityUsecase(
	skillRepo *repoSkill.Queries,
	userRepo *repoUser.Queries,
	activityRepo *repoActivity.Queries,
) ActivityUsecase {
	return &activityUsecase{
		skillRepo:    skillRepo,
		userRepo:     userRepo,
		activityRepo: activityRepo,
	}
}

type activityUsecase struct {
	skillRepo    *repoSkill.Queries
	userRepo     *repoUser.Queries
	activityRepo *repoActivity.Queries
}
