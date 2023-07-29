package ucActivity

import (
	"context"

	dtoActivity "github.com/kbiits/gikslab_assesment/internal/app/dto/activity"
	repoActivity "github.com/kbiits/gikslab_assesment/internal/app/repository/activity"
	repoSkill "github.com/kbiits/gikslab_assesment/internal/app/repository/skill"
	repoUser "github.com/kbiits/gikslab_assesment/internal/app/repository/user"
)

func (uc *activityUsecase) ListActivityBySkillId(ctx context.Context, req *dtoActivity.ListActivitiesDtoRequest) (resp dtoActivity.ListActivitiesDtoResponse, err error) {
	offset := req.Limit * (req.Page - 1)
	queryActivitiesParams := repoActivity.ListActivitiesBySkillParams{
		Skillid:   req.SkillId,
		Limitval:  int32(req.Limit),
		Offsetval: int32(offset),
		Sortby:    *req.SortBy,
		Ordercol:  *req.Order,
	}
	activities, err := uc.activityRepo.ListActivitiesBySkill(ctx, queryActivitiesParams)
	if err != nil {
		return
	}

	mapSkillIds := make(map[int64]bool)
	mapParticipantIds := make(map[int64]bool)
	for _, a := range activities {
		mapSkillIds[a.SkillID] = true

		participantsArr := make([]int64, 0)
		err = json.Unmarshal([]byte(a.Participants), &participantsArr)
		if err != nil {
			return
		}

		for _, v := range participantsArr {
			mapParticipantIds[v] = true
		}
	}

	skillIds := make([]int64, 0)
	participantIds := make([]int64, 0)

	for id := range mapSkillIds {
		skillIds = append(skillIds, id)
	}
	for id := range mapParticipantIds {
		participantIds = append(participantIds, id)
	}

	users, err := uc.userRepo.FindUsersByMultipleIds(ctx, participantIds)
	if err != nil {
		return
	}

	skills, err := uc.skillRepo.FindSkillByMultipleIds(ctx, skillIds)
	if err != nil {
		return
	}

	mapSkills := make(map[int64]*repoSkill.Skill)

	for i := 0; i < len(skills); i++ {
		s := skills[i]
		mapSkills[s.ID] = &s
	}

	mapUsers := make(map[int64]*repoUser.User)
	skillIdsForUsers := make([]int64, 0)
	for i := 0; i < len(users); i++ {
		u := users[i]
		mapUsers[u.ID] = &u

		if u.SkillID.Valid {
			skillIdsForUsers = append(skillIdsForUsers, u.SkillID.Int64)
		}
	}

	skillsForUsers, err := uc.skillRepo.FindSkillByMultipleIds(ctx, skillIdsForUsers)
	if err != nil {
		return
	}

	mapSkillsForUsers := make(map[int64]*repoSkill.Skill)
	for i := 0; i < len(skillsForUsers); i++ {
		skill := skillsForUsers[i]
		mapSkillsForUsers[skill.ID] = &skill
	}

	resp = make(dtoActivity.ListActivitiesDtoResponse, 0)
	for _, a := range activities {
		var act dtoActivity.ActivityDto

		skill, ok := mapSkills[a.SkillID]
		if skill != nil && ok {
			act.SkillId = skill.ID
			act.SkillName = skill.Name
		}

		act.Id = a.ID
		act.Title = a.Title
		act.Description = a.Description
		act.StartDate = a.StartDate.Format(dtoActivity.ActivityDateLayout)
		act.EndDate = a.EndDate.Format(dtoActivity.ActivityDateLayout)

		act.Participants = make([]dtoActivity.ParticipantDto, 0)

		users := make([]int64, 0)

		err = json.Unmarshal([]byte(a.Participants), &users)
		if err != nil {
			return
		}

		for _, actUser := range users {
			u, ok := mapUsers[actUser]
			if ok && u != nil {
				skill := mapSkillsForUsers[u.SkillID.Int64]
				p := dtoActivity.ParticipantDto{
					Id:      u.ID,
					Name:    u.Name,
					Profile: u.Profile,
				}

				if skill != nil {
					p.Skill = skill.Name
				}

				act.Participants = append(act.Participants, p)
			}
		}

		resp = append(resp, act)
	}

	return
}
