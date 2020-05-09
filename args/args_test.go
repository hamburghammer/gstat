package args_test

import (
	"testing"

	"github.com/hamburghammer/gstat/args"
)

func TestValidate(t *testing.T) {
	baseError := e.baseError
	t.Run("should error if args are empty", func(t *testing.T) {
		arguments := args.Arguments{}

		got := arguments.Validate()
		want := args.ValidationError{}
	})
}
