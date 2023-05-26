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

type AuthenticationError struct {
	Err         error
	Message     string
	UserMessage string
	Stack       []string
}

func createError() *AuthenticationError {
	return &AuthenticationError{
		Stack: []string{},
	}
}

func (a AuthenticationError) New() Error {
	newError := createError()
	newError.Message = a.Message
	return newError
}

func (a AuthenticationError) Wrap(err error) Error {
	newError := createError()
	a.Err = err
	a.Message = fmt.Sprintf("%s; %s", a.Message, err.Error())
	return newError
}

func (a *AuthenticationError) WrapWithUserMessage(err error, userMessage string, args ...any) Error {
	newError := createError()
	a.Err = err
	a.Message = fmt.Sprintf("%s; %s", a.Message, err.Error())
	a.UserMessage = fmt.Sprintf(userMessage, args)
	return newError
}

func (a *AuthenticationError) Error() string {
	return a.Message
}

func (a *AuthenticationError) With(key string, val interface{}) Error {
	a.Message = fmt.Sprintf("%s; [%s]=%s", a.Message, key, val)
	return a
}

func (a AuthenticationError) Log() {
	log.Println(a.Message)
}

func (a AuthenticationError) AddUserMessage(formattedMessage string, params ...any) {
	a.UserMessage = fmt.Sprintf(formattedMessage, params)
}

func (a *AuthenticationError) JSON() map[string]any {
	if a.UserMessage == "" {
		a.UserMessage = "Failed to perform request"
	}

	return map[string]any{
		"errorMessage": a.Message,
		"userMessage":  a.UserMessage,
	}
}
