package error_handler

import (
	"github.com/sirupsen/logrus"
)

type CustomError struct {
	Code    int
	Message string
	Err     error
}

func (ce *CustomError) Error() string {
	return ce.Message
}

func NewCustomError(code int, newMsg string, newErr error) *CustomError {
	logrus.WithFields(logrus.Fields{
		"code":    code,
		"message": newMsg,
		"error":   newErr,
	})
	return &CustomError{Code: code, Message: newMsg, Err: newErr}
}
