package dtoCommon

type ResponseWithMessage struct {
	Message string `json:"message"`
}

func NewResponseWithMessage(message string) ResponseWithMessage {
	return ResponseWithMessage{
		Message: message,
	}
}
