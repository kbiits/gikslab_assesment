package errors

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kbiits/gikslab_assesment/internal/app/constants"
	dtoCommon "github.com/kbiits/gikslab_assesment/internal/app/dto/common"
)

var HttpErrorHandler = func(c *fiber.Ctx, err error) error {

	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	var httpError *HttpError
	var e *fiber.Error

	requestId := c.Locals(constants.KeyRequestId)
	path := c.Path()

	if errors.As(err, &httpError) {
		code = httpError.StatusCode
	} else if errors.As(err, &e) {
		// Retrieve the custom status code if it's a *fiber.Error
		log.Printf("fiber error, err : %#v\n", e)
		code = e.Code
	} else {
		log.Printf("unknown error, err : %#v\n", err)
		err = errors.New("internal server error")
	}

	log.Printf("`%v` REQ ID : %v - Caught ERROR, code: %v, message: %v\n", path, requestId, code, err.Error())

	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Return status code with error message
	return c.Status(code).JSON(dtoCommon.NewResponseWithMessage(err.Error()))
}
