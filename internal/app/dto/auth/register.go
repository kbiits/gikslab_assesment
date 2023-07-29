package dtoAuth

type RegisterDtoRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
	Skill    string `json:"skill"`
}

type RegisterDtoResponse struct {
	Message string `json:"message"`
}

func NewRegisterDtoResponseSuccess() RegisterDtoResponse {
	return RegisterDtoResponse{
		Message: "create success",
	}
}
