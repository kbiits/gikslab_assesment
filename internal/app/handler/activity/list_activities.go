package handlerActivity

import (
	"log"

	"github.com/gofiber/fiber/v2"
	dtoActivity "github.com/kbiits/gikslab_assesment/internal/app/dto/activity"
	errorActivity "github.com/kbiits/gikslab_assesment/internal/app/errors/activity"
)

func (h *activityHandler) ListActivities(ctx *fiber.Ctx) (err error) {

	req, err := dtoActivity.NewListActivitiesRequest(ctx)
	if err != nil {
		return
	}

	resp, err := h.activityUc.ListActivityBySkillId(ctx.Context(), &req)
	if err != nil {
		log.Println(err)
		return errorActivity.ErrListActivity
	}

	return ctx.JSON(resp)
}
