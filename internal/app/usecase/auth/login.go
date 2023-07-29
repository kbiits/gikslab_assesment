package ucAuth

import (
	"context"
	"database/sql"
	"log"

	dtoAuth "github.com/kbiits/gikslab_assesment/internal/app/dto/auth"
	errorAuth "github.com/kbiits/gikslab_assesment/internal/app/errors/auth"
	"github.com/kbiits/gikslab_assesment/internal/app/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (uc *authUsecase) Login(ctx context.Context, req *dtoAuth.LoginDtoRequest) (resp dtoAuth.LoginDtoResponse, err error) {
	user, err := uc.userRepo.FindUserByUsername(ctx, req.Username)
	if err == sql.ErrNoRows {
		err = errorAuth.ErrLoginInvalid
		return
	}

	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Println(err)
		return resp, errorAuth.ErrLoginInvalid
	}

	token, err := uc.jwt.CreateToken(jwt.Claims{
		UserId:   int(user.ID),
		Username: user.Username,
		Profile:  user.Profile,
	})
	if err != nil {
		return
	}

	resp = dtoAuth.NewLoginDtoResponse(token, user.Profile)
	return
}
