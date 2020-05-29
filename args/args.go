package args

import (
	"errors"
	"strings"

	"github.com/jessevdk/go-flags"
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
func (a *Arguments) Validate() []ValidationError {
	validationErrors := make([]ValidationError, 0, 10)
	validationErrors = append(validationErrors, newValidationError(*a, "Arguments are empty"))
	return validationErrors
}

// Parse the flags to the Arguments struct.
func Parse() Arguments {

	args := Arguments{}

	_, err := flags.Parse(&args)

	if err != nil {
		panic(err)
	}

	return args
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
