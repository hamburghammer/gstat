package args

import (
	"errors"
	"fmt"
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

// Equals checks for field equality
func (a Arguments) Equals(other Arguments) bool {
	if a.CPU != other.CPU {
		return false
	}
	if a.Disk != other.Disk {
		return false
	}
	if a.Mem != other.Mem {
		return false
	}
	if a.Processes != other.Processes {
		return false
	}
	if a.Health != other.Health {
		return false
	}
	return true
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
