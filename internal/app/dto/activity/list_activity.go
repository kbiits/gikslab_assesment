package dtoActivity

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	errorActivity "github.com/kbiits/gikslab_assesment/internal/app/errors/activity"
)

type ListActivitiesDtoRequest struct {
	SortBy  *string
	Order   *string
	Page    int
	Limit   int
	SkillId int64
}

var availableSortBy = map[string]bool{
	"skill_id":    true,
	"title":       true,
	"description": true,
	"start_date":  true,
	"end_date":    true,
}

var mapOrder = map[string]string{
	"asc":        "ASC",
	"ascending":  "ASC",
	"desc":       "DESC",
	"descending": "DESC",
}

func NewListActivitiesRequest(ctx *fiber.Ctx) (req ListActivitiesDtoRequest, err error) {
	sortBy := ctx.Query("sort_by")
	if _, ok := availableSortBy[sortBy]; sortBy != "" && !ok {
		return req, errorActivity.ErrListActivity
	} else if ok {
		req.SortBy = &sortBy
	} else {
		sortBy := "start_date"
		req.SortBy = &sortBy
	}

	order := ctx.Query("order")
	if v, ok := mapOrder[order]; order != "" && !ok {
		return req, errorActivity.ErrListActivity
	} else if v != "" {
		req.Order = &v
	} else {
		order := "ASC"
		req.Order = &order
	}

	pageStr := ctx.Query("page")
	if pageStr == "" {
		req.Page = 1
	} else {
		pageInt, err := strconv.Atoi(pageStr)
		if err != nil || pageInt == 0 {
			return req, errorActivity.ErrListActivity
		}
		req.Page = pageInt
	}

	limitStr := ctx.Query("limit")
	if limitStr == "" {
		req.Limit = 10
	} else {
		limitInt, err := strconv.Atoi(limitStr)
		if err != nil || limitInt == 0 {
			return req, errorActivity.ErrListActivity
		}
		req.Limit = limitInt
	}

	skillId := ctx.Params("skillId")
	if skillId == "" {
		return req, errorActivity.ErrListActivity
	} else {
		skillIdInt, err := strconv.ParseInt(skillId, 10, 64)
		if err != nil {
			return req, errorActivity.ErrListActivity
		}
		req.SkillId = skillIdInt
	}

	return
}

type ListActivitiesDtoResponse []ActivityDto

type ActivityDto struct {
	Id           int64            `json:"id"`
	SkillId      int64            `json:"skill_id"`
	SkillName    string           `json:"skill_name"`
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	StartDate    string           `json:"startdate"`
	EndDate      string           `json:"enddate"`
	Participants []ParticipantDto `json:"participants"`
}

type ParticipantDto struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
	Skill   string `json:"skill"`
}
