package web

import "net/http"

type AppError struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewAppError(message string, code int, causes []Cause) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
		Causes:  causes,
	}
}

func UnprocessableEntityError(message string) *AppError {
	return NewAppError(message, http.StatusUnprocessableEntity, nil)
}

func BadRequestErrorWithCauses(message string, causes []Cause) *AppError {
	return NewAppError(message, http.StatusBadRequest, causes)
}

func ConflictError(message string) *AppError {
	return NewAppError(message, http.StatusConflict, nil)
}
