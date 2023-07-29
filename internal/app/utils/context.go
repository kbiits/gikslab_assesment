package utils

import (
	"context"
	"errors"

	"github.com/kbiits/gikslab_assesment/internal/app/constants"
	"github.com/kbiits/gikslab_assesment/internal/app/jwt"
)

func GetLoggedInUserFromContext(ctx context.Context) (claims jwt.Claims, err error) {
	claims, ok := ctx.Value(constants.KeyLoggedInUser).(jwt.Claims)
	if !ok {
		err = errors.New("cant convert to jwt claims")
	}

	return
}
