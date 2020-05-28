package args_test

import (
	"testing"

	"github.com/hamburghammer/gstat/args"
	e "github.com/hamburghammer/gstat/errors"
)

func TestValidate(t *testing.T) {
	baseError := e.BaseError{Operation: "Validation"}

	t.Run("should error if args are empty", func(t *testing.T) {
		arguments := args.Arguments{}
		baseError.Message = "Arguments are empty"

		got := arguments.Validate()
		want := args.ValidationError{BaseError: baseError, Arguments: arguments}

		if len(got) != 1 {
			t.Errorf("An error was expected but got: %v", got)
		}

		gotValitaionError := got[0]

		if gotValitaionError.Message != want.Message {
			t.Errorf("got error message: %s \n want message: %s", gotValitaionError.Message, want.Message)
		}

	})
}
