package args

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jessevdk/go-flags"

	e "github.com/hamburghammer/gstat/errors"
)

// Arguments represent the flags given at program start.
type Arguments struct {
	CPU       bool   `short:"c" long:"cpu" description:"Include the total CPU consumption"`
	Mem       bool   `short:"m" long:"mem" description:"Include the total RAM consumption"`
	Disk      bool   `short:"d" long:"disk" description:"Include the total CPU consumption"`
	Processes bool   `short:"p" long:"proc" description:"Include the top 10 processes"`
	Health    string `long:"health" description:"Make a healthcheck call against the URI"`
	rest      []string
}

// Validate the arguments.
func (a *Arguments) Validate() []error {
	validationErrors := make([]error, 1)
	return append(validationErrors, errors.New("Validate not impmented jet"))
}

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

// Parse the flags to the Arguments struct.
func Parse() Arguments {

	args := Arguments{}

	re, err := flags.Parse(&args)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Return value from parsing args: %v \n", re)
	return args
}

func newValidationError(args Arguments, message string) ValidationError {
	return ValidationError{e.BaseError{Operation: "Validation", Message: message}, args}
}

func uriValidate(uri string) error {
	scheme := strings.Split(uri, "://")

	if len(scheme) < 2 {
		return errors.New("The URI does not looks like schema://provider")
	}

	if len(strings.Split(scheme[1], ".")) < 2 {
		return errors.New("The URI provider does not has a top and second level domain like example.com")
	}

	return nil
}
