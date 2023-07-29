package ucActivity

import (
	"context"
	"database/sql"
	"log"

	jsoniter "github.com/json-iterator/go"
	dtoActivity "github.com/kbiits/gikslab_assesment/internal/app/dto/activity"
	errorActivity "github.com/kbiits/gikslab_assesment/internal/app/errors/activity"
	repoActivity "github.com/kbiits/gikslab_assesment/internal/app/repository/activity"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (uc *activityUsecase) AddActivity(ctx context.Context, req *dtoActivity.AddActivityDtoRequest) (resp dtoActivity.AddActivityDtoResponse, err error) {
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
		log.Println("there's not found user")
		return resp, errorActivity.ErrFailedAddActivity
	}

	userIds := make([]int64, 0)
	for _, user := range users {
		if user.SkillID.Int64 != skill.ID {
			return resp, errorActivity.ErrFailedAddActivity
		}

		userIds = append(userIds, user.ID)
	}

	participantsBytes, err := json.Marshal(userIds)
	if err != nil {
		return
	}

	paramsAdd := repoActivity.AddActivityParams{
		Title:        req.Title,
		Description:  req.Description,
		StartDate:    req.StartDateParsed,
		EndDate:      req.EndDateParsed,
		SkillID:      skill.ID,
		Participants: string(participantsBytes),
	}
	err = uc.activityRepo.AddActivity(ctx, paramsAdd)
	if err != nil {
		return
	}

	resp = dtoActivity.NewAddActivityDtoResponse()
	return
}
