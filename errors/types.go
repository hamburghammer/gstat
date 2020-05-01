package errors

import "fmt"

type BaseError struct {
	Operation string
	Message string
}

func (e *BaseError) Error() string {
	return fmt.Sprintf("%s failed because of %s", e.Operation, e.Message)
}