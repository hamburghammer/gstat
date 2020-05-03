package args

import (
	"errors"
	"fmt"

	"github.com/jessevdk/go-flags"

	e "github.com/hamburghammer/gstat/errors"
)

type Arguments struct {
	Cpu       bool   `short:"c" long:"cpu" description:"Include the total CPU consumption"`
	Mem       bool   `short:"m" long:"mem" description:"Include the total RAM consumption"`
	Disk      bool   `short:"d" long:"disk" description:"Include the total CPU consumption"`
	Processes bool   `short:"p" long:"proc" description:"Include the top 10 processes"`
	Health    string `long:"health" description:"Make a healthcheck call"`
	rest      []string
}

type ValidationError struct {
	e.BaseError
	Argument string
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("%s of the argument %s failed: %s", ve.Operation, ve.Argument, ve.Message)
}

func (a *Arguments) Validate() error {
	return errors.New("Validate not impmented jet")
}

func Parse() Arguments {

	args := Arguments{}

	re, err := flags.Parse(&args)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Return value from parsing args: %v \n", re)
	return args
}
