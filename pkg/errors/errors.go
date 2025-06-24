package errors

import "fmt"

type ErrorType string

const (
	InternalError ErrorType = "INTERNAL"
	BusinessError ErrorType = "BUSINESS"
)

type AppError struct {
	Type    ErrorType
	Message string
	Code    int // optional: http status code
	Err     error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s: %v", e.Type, e.Message, e.Err)
}

func NewInternalError(msg string, err error) *AppError {
	return &AppError{
		Type:    InternalError,
		Message: msg,
		Err:     err,
		Code:    500,
	}
}

func NewBusinessError(msg string) *AppError {
	return &AppError{
		Type:    BusinessError,
		Message: msg,
		Code:    400,
	}
}
