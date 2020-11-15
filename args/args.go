package args

import (
	"errors"
	"fmt"
	"os"
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
	var validationErrors = make([]ValidationError, 0, 10)

	if a.Equals(Arguments{}) {
		return append(validationErrors, newValidationError(*a, "Arguments are empty"))
	}

	err := uriValidate(a.Health)
	if err != nil {
		validationErrors = append(validationErrors, newValidationError(*a, err.Error()))
	}

	return validationErrors
}

// Equals checks for field equality
func (a Arguments) Equals(other Arguments) bool {
	if a.CPU != other.CPU {
		return false
	}
	if a.Mem != other.Mem {
		return false
	}
	if a.Disk != other.Disk {
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

	_, err := flags.Parse(&args)

	if err != nil {
		if _, ok := err.(*flags.Error); ok {
			os.Exit(1)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
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
