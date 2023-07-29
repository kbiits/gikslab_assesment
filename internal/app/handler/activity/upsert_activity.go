package handlerActivity

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	dtoActivity "github.com/kbiits/gikslab_assesment/internal/app/dto/activity"
	errorActivity "github.com/kbiits/gikslab_assesment/internal/app/errors/activity"
)

func (h *activityHandler) UpsertActivity(ctx *fiber.Ctx) (err error) {
	req := new(dtoActivity.AddActivityDtoRequest)
	err = ctx.BodyParser(req)
	if err != nil {
		if err == fiber.ErrUnprocessableEntity {
			err = errorActivity.ErrFailedAddActivity
		}

		return
	}

	err = req.ParseAndValidateRequest()
	if err != nil {
		return
	}

	actId := ctx.Params("activityId")
	var resp interface{}
	if actId == "" {
		resp, err = h.activityUc.AddActivity(ctx.Context(), req)
		if err != nil {
			return errorActivity.ErrFailedAddActivity
		}
	} else {
		actIdInt, err := strconv.Atoi(actId)
		if err != nil {
			return errorActivity.ErrFailedUpdateActivity
		}
		resp, err = h.activityUc.UpdateActivity(ctx.Context(), req, actIdInt)
		if err != nil {
			return errorActivity.ErrFailedUpdateActivity
		}
	}

	return ctx.JSON(resp)
}
