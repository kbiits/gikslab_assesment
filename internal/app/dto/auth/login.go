package dtoAuth

type LoginDtoRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginDtoResponse struct {
	Token   string `json:"token"`
	Profile string `json:"profile"`
}

func NewLoginDtoResponse(token, profile string) LoginDtoResponse {
	return LoginDtoResponse{
		Token:   token,
		Profile: profile,
	}
}
