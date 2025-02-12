package Utility

import (
	"fmt"
)

type ErrorHandler struct {
	errors []string
}

func (handler *ErrorHandler) AddError(fmtMessage string, args ...interface{}) {
	err := fmt.Sprintf(fmtMessage, args...)
	handler.errors = append(handler.errors, err)
}

func (handler *ErrorHandler) AddErrors(errors []string, lastError string) {
	for _, err := range errors {
		handler.AddError(err)
	}

	handler.AddError(lastError)
}

func (handler *ErrorHandler) GetErrors() []string {
	return handler.errors
}
