package warrant

import "fmt"

type Error struct {
	Message      string `json:"message"`
	WrappedError error  `json:"-"`
}

func (err Error) Error() string {
	if err.WrappedError != nil {
		return fmt.Sprintf("Warrant error: %s %s", err.Message, err.WrappedError.Error())
	}
	return fmt.Sprintf("Warrant error: %s", err.Message)
}

func WrapError(message string, err error) Error {
	return Error{
		Message:      message,
		WrappedError: err,
	}
}
