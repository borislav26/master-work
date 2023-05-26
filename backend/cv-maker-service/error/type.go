package error

import (
	"fmt"
	"log"
)

type Error interface {
	Error() string
	With(key string, val interface{}) Error
	Log()
	JSON() map[string]any
}

type CVMakerError struct {
	Err         error
	Message     string
	UserMessage string
	Stack       []string
}

func createError() *CVMakerError {
	return &CVMakerError{
		Stack: []string{},
	}
}

func (a CVMakerError) New() Error {
	newError := createError()
	newError.Message = a.Message
	return newError
}

func (a CVMakerError) Wrap(err error) Error {
	newError := createError()
	a.Err = err
	a.Message = fmt.Sprintf("%s; %s", a.Message, err.Error())
	return newError
}

func (a *CVMakerError) WrapWithUserMessage(err error, userMessage string, args ...any) Error {
	newError := createError()
	newError.Err = err
	newError.Message = fmt.Sprintf("%s; %s", a.Message, err.Error())
	newError.UserMessage = fmt.Sprintf(userMessage, args)
	return newError
}

func (a *CVMakerError) Error() string {
	return a.Message
}

func (a *CVMakerError) With(key string, val interface{}) Error {
	a.Message = fmt.Sprintf("%s; [%s]=%s", a.Message, key, val)
	return a
}

func (a CVMakerError) Log() {
	log.Println(a.Message)
}

func (a CVMakerError) AddUserMessage(formattedMessage string, params ...any) {
	a.UserMessage = fmt.Sprintf(formattedMessage, params)
}

func (a *CVMakerError) JSON() map[string]any {
	if a.UserMessage == "" {
		a.UserMessage = "Failed to perform request"
	}

	return map[string]any{
		"errorMessage": a.Message,
		"userMessage":  a.UserMessage,
	}
}
