package ucActivity

import (
	"context"

	dtoActivity "github.com/kbiits/gikslab_assesment/internal/app/dto/activity"
	errorActivity "github.com/kbiits/gikslab_assesment/internal/app/errors/activity"
)

func (uc *activityUsecase) DeleteActivity(ctx context.Context, actId int) (resp dtoActivity.DeleteActivityDtoResponse, err error) {
	affected, err := uc.activityRepo.DeleteActivity(ctx, int64(actId))
	if err != nil || affected == 0 {
		err = errorActivity.ErrFailedDeleteActivity
		return
	}

	resp = dtoActivity.NewDeleteActivityDtoResponse()
	return
}
