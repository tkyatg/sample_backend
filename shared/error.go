package shared

type ResponseError struct {
	Message string `json:"message"`
}

func NewResponseError(errMessage string) ResponseError {
	return ResponseError{
		Message: errMessage,
	}
}
