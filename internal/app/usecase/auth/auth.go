package ucAuth

import (
	"context"

	dtoAuth "github.com/kbiits/gikslab_assesment/internal/app/dto/auth"
	"github.com/kbiits/gikslab_assesment/internal/app/jwt"
	"github.com/kbiits/gikslab_assesment/internal/app/repository/redis"
	repoSkill "github.com/kbiits/gikslab_assesment/internal/app/repository/skill"
	repoUser "github.com/kbiits/gikslab_assesment/internal/app/repository/user"
)

type AuthUsecase interface {
	Login(ctx context.Context, req *dtoAuth.LoginDtoRequest) (dtoAuth.LoginDtoResponse, error)
	Logout(ctx context.Context, token string) (dtoAuth.LogoutDtoResponse, error)
	Register(ctx context.Context, req *dtoAuth.RegisterDtoRequest) (dtoAuth.RegisterDtoResponse, error)
}

func NewAuthUsecase(
	userRepo *repoUser.Queries,
	skillRepo *repoSkill.Queries,
	jwt *jwt.Jwt,
	redisRepo redis.RedisRepository,
) AuthUsecase {
	return &authUsecase{
		userRepo:  userRepo,
		skillRepo: skillRepo,
		jwt:       jwt,
		redisRepo: redisRepo,
	}
}

type authUsecase struct {
	userRepo  *repoUser.Queries
	skillRepo *repoSkill.Queries
	jwt       *jwt.Jwt
	redisRepo redis.RedisRepository
}
