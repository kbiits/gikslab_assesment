package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/kbiits/gikslab_assesment/internal/app/constants"
)

type Claims struct {
	UserId    int
	Username  string
	Profile   string
	ExpiredAt int64
	IssuedAt  int64
}

func (c *Claims) IsProfileBoard() bool {
	return c.Profile == constants.ProfileBoard
}

func (c *Claims) IsProfileExpert() bool {
	return c.Profile == constants.ProfileExpert
}

func (c Claims) ToJwtMapClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"username": c.Username,
		"userId":   c.UserId,
		"profile":  c.Profile,
		"iat":      c.IssuedAt,
		"exp":      c.ExpiredAt,
	}
}
