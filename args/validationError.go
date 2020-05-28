package args

import (
	"fmt"

	e "github.com/hamburghammer/gstat/errors"
)

// OperationKeyValidation represents the key for the Operation field of an ValidationError
const OperationKeyValidation = "Validation"

// ValidationError is a struct to wrap the error with more information.
type ValidationError struct {
	e.BaseError
	Arguments
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("%s of the arguments %+v failed: %s", ve.Operation, ve.Arguments, ve.Message)
}

func newValidationError(args Arguments, message string) ValidationError {
	return ValidationError{e.BaseError{Operation: "Validation", Message: message}, args}
}
