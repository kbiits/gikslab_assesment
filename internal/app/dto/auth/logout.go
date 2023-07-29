package dtoAuth

type LogoutDtoResponse struct {
	Message string `json:"message"`
}

func NewLogoutDtoResponse() LogoutDtoResponse {
	return LogoutDtoResponse{
		Message: "logout success",
	}
}
