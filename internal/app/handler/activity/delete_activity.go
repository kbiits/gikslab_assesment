package handlerActivity

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	errorActivity "github.com/kbiits/gikslab_assesment/internal/app/errors/activity"
)

func (h *activityHandler) DeleteActivity(ctx *fiber.Ctx) (err error) {
	actId := ctx.Params("activityId")

	if actId == "" {
		return errorActivity.ErrFailedDeleteActivity
	}

	actIdInt, err := strconv.Atoi(actId)
	if err != nil {
		return errorActivity.ErrFailedDeleteActivity
	}

	resp, err := h.activityUc.DeleteActivity(ctx.Context(), actIdInt)
	if err != nil {
		return errorActivity.ErrFailedDeleteActivity
	}

	return ctx.JSON(resp)
}
