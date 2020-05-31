package args_test

import (
	"testing"

	"github.com/hamburghammer/gstat/args"
)

func TestValidate(t *testing.T) {

	t.Run("should error if args are empty", func(t *testing.T) {
		arguments := args.Arguments{}

		got := arguments.Validate()
		want := "Arguments are empty"

		if len(got) != 1 {
			t.Errorf("An error was expected but got '%d' errors", len(got))
		}

		gotValidationError := got[0]

		if gotValidationError.Message != want {
			t.Errorf("got error message: '%s' \n want message: '%s'", gotValidationError.Message, want)
		}

	})

	t.Run("should have url validation error", func(t *testing.T) {
		arguments := args.Arguments{}
		arguments.Health = "example.com"

		got := arguments.Validate()
		want := "The URI does not looks like schema://provider"

		if len(got) != 1 {
			t.Errorf("An error was expected but got '%d' errors", len(got))
		}

		gotValidationError := got[0]

		if gotValidationError.Message != want {
			t.Errorf("got error message: '%s' \n want message: '%s'", gotValidationError.Message, want)
		}

	})
}

func TestEquals(t *testing.T) {
	defaultArgs := args.Arguments{}

	t.Run("should check CPU", func(t *testing.T) {
		args := args.Arguments{CPU: true}

		if args.Equals(defaultArgs) {
			t.Fail()
		}

	})

	t.Run("should check Mem", func(t *testing.T) {
		args := args.Arguments{Mem: true}

		if args.Equals(defaultArgs) {
			t.Fail()
		}

	})

	t.Run("should check Disk", func(t *testing.T) {
		args := args.Arguments{Disk: true}

		if args.Equals(defaultArgs) {
			t.Fail()
		}

	})

	t.Run("should check Processes", func(t *testing.T) {
		args := args.Arguments{Processes: true}

		if args.Equals(defaultArgs) {
			t.Fail()
		}

	})

	t.Run("should check Health", func(t *testing.T) {
		args := args.Arguments{Health: "http://example.com"}

		if args.Equals(defaultArgs) {
			t.Fail()
		}

	})
}
