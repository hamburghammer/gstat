package args_test

import (
	"testing"

	"github.com/hamburghammer/gstat/args"
	e "github.com/hamburghammer/gstat/errors"
)

func TestErrorFormatting(t *testing.T) {

	validationError := args.ValidationError{BaseError: e.BaseError{Operation: "Validation", Message: "Test message"}, Arguments: args.Arguments{}}
	got := validationError.Error()
	want := "Validation of the arguments {CPU:false Mem:false Disk:false Processes:false Health: rest:[]} failed: Test message"

	if got != want {
		t.Errorf("Want: %s but got: %s", want, got)
	}
}
