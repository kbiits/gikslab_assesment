package ucAuth

import (
	"context"
	"fmt"
	"log"
	"time"

	dtoAuth "github.com/kbiits/gikslab_assesment/internal/app/dto/auth"
	"github.com/kbiits/gikslab_assesment/internal/app/errors"
	"github.com/kbiits/gikslab_assesment/internal/app/jwt"
	"github.com/kbiits/gikslab_assesment/internal/app/utils"
)

func (uc *authUsecase) Logout(ctx context.Context, token string) (resp dtoAuth.LogoutDtoResponse, err error) {
	userClaims, err := utils.GetLoggedInUserFromContext(ctx)
	if err != nil {
		log.Println(err)
		return resp, errors.ErrUnauthorized
	}

	now := time.Now()
	expiredAt := time.Unix(userClaims.ExpiredAt, 0)
	durationExpired := expiredAt.Sub(now)

	err = uc.redisRepo.Set(ctx, fmt.Sprintf(jwt.FormatRedisKeyBlacklisted, token), "logged_out", durationExpired)
	if err != nil {
		return
	}

	resp = dtoAuth.NewLogoutDtoResponse()
	return
}
