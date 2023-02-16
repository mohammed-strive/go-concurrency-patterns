package errors

import "fmt"

type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func (err MyError) Error() string {
	return err.Message
}

func NewError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError{
		Inner:   err,
		Message: fmt.Sprintf(messagef, msgArgs...),
		Misc:    make(map[string]interface{}),
	}
}
