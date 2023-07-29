package ucActivity

import (
	"context"
	"database/sql"
	"log"

	dtoActivity "github.com/kbiits/gikslab_assesment/internal/app/dto/activity"
	errorActivity "github.com/kbiits/gikslab_assesment/internal/app/errors/activity"
	repoActivity "github.com/kbiits/gikslab_assesment/internal/app/repository/activity"
)

func (uc *activityUsecase) UpdateActivity(ctx context.Context, req *dtoActivity.AddActivityDtoRequest, actId int) (resp dtoActivity.UpdateActivityDtoResponse, err error) {
	act, err := uc.activityRepo.FindActivityById(ctx, int64(actId))
	if err != nil {
		err = errorActivity.ErrFailedUpdateActivity
		return
	}

	skill, err := uc.skillRepo.FindSkillByName(ctx, req.Skill)
	if err == sql.ErrNoRows {
		err = errorActivity.ErrFailedAddActivity
		return
	}

	if err != nil {
		return
	}

	participantsMap := make(map[string]bool)
	for _, v := range req.Participants {
		participantsMap[v] = true
	}

	participantsArray := make([]string, 0)
	for p := range participantsMap {
		participantsArray = append(participantsArray, p)
	}

	users, err := uc.userRepo.FindUsersByUsernames(ctx, participantsArray)
	if err != nil {
		return
	}

	if len(users) != len(participantsArray) {
		log.Println("there's not found user in participants")
		return resp, errorActivity.ErrFailedAddActivity
	}

	userIds := make([]int64, 0)
	for _, user := range users {
		if user.SkillID.Int64 != skill.ID {
			return resp, errorActivity.ErrFailedUpdateActivity
		}

		userIds = append(userIds, user.ID)
	}

	participantsBytes, err := json.Marshal(userIds)
	if err != nil {
		return
	}

	paramsUpdate := repoActivity.UpdateActivityParams{
		Title:        req.Title,
		Description:  req.Description,
		StartDate:    req.StartDateParsed,
		EndDate:      req.EndDateParsed,
		SkillID:      skill.ID,
		Participants: string(participantsBytes),
		ID:           act.ID,
	}
	err = uc.activityRepo.UpdateActivity(ctx, paramsUpdate)
	if err != nil {
		return
	}

	resp = dtoActivity.NewUpdateActivityDtoResponse()
	return
}
