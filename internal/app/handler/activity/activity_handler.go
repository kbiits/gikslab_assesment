package handlerActivity

import (
	"github.com/gofiber/fiber/v2"
	ucActivity "github.com/kbiits/gikslab_assesment/internal/app/usecase/activity"
)

type activityHandler struct {
	activityUc ucActivity.ActivityUsecase
}

type ActivityHandler interface {
	ListActivities(ctx *fiber.Ctx) error
	UpsertActivity(ctx *fiber.Ctx) error
	DeleteActivity(ctx *fiber.Ctx) error
}

func NewActivityHandler(activityUc ucActivity.ActivityUsecase) ActivityHandler {
	return &activityHandler{
		activityUc: activityUc,
	}
}
