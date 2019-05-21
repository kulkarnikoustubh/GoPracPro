package errors

import (
	"fmt"
)

const (
	errMsgPrefix = "Rest_Client_Error : "
)

//RestClientError hold error for rest client having customized error message
type RestClientError interface {
	error
}

type restClientErrorImpl struct {
	msg string
}

func (e restClientErrorImpl) Error() string { return e.msg }

// New returns a new RestClientError
func New(message string) RestClientError {
	return restClientErrorImpl{
		msg: errMsgPrefix + message,
	}
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorf(format string, args ...interface{}) RestClientError {
	return restClientErrorImpl{
		msg: errMsgPrefix + fmt.Sprintf(format, args...),
	}
}
