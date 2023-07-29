package dtoActivity

import (
	"log"
	"time"

	errorActivity "github.com/kbiits/gikslab_assesment/internal/app/errors/activity"
)

type AddActivityDtoRequest struct {
	Skill           string   `json:"skill"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	StartDate       string   `json:"startdate"`
	EndDate         string   `json:"enddate"`
	Participants    []string `json:"participants"`
	StartDateParsed time.Time
	EndDateParsed   time.Time
}

const ActivityDateLayout = "2006-01-02 15:04:05"

func (r *AddActivityDtoRequest) ParseAndValidateRequest() (err error) {
	if r.Skill == "" || r.Title == "" || r.Description == "" || r.StartDate == "" || r.EndDate == "" || len(r.Participants) == 0 {
		return errorActivity.ErrFailedAddActivity
	}

	r.StartDateParsed, err = time.Parse(ActivityDateLayout, r.StartDate)
	if err != nil {
		log.Println(err)
		return errorActivity.ErrFailedAddActivity
	}
	r.EndDateParsed, err = time.Parse(ActivityDateLayout, r.EndDate)
	if err != nil {
		return errorActivity.ErrFailedAddActivity
	}

	if r.EndDateParsed.Before(r.StartDateParsed) {
		return errorActivity.ErrFailedAddActivity
	}

	return nil
}

type AddActivityDtoResponse struct {
	Message string `json:"message"`
}

func NewAddActivityDtoResponse() AddActivityDtoResponse {
	return AddActivityDtoResponse{
		Message: "create success",
	}
}

type UpdateActivityDtoResponse struct {
	Message string `json:"message"`
}

func NewUpdateActivityDtoResponse() UpdateActivityDtoResponse {
	return UpdateActivityDtoResponse{
		Message: "update success",
	}
}

type DeleteActivityDtoResponse struct {
	Message string `json:"message"`
}

func NewDeleteActivityDtoResponse() DeleteActivityDtoResponse {
	return DeleteActivityDtoResponse{
		Message: "delete success",
	}
}
