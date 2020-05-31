package errors_test

import (
	"testing"

	"github.com/hamburghammer/gstat/errors"
)

func TestBaseError(t *testing.T) {
	baseError := errors.BaseError{Operation: "Testing", Message: "Test message"}

	got := baseError.Error()
	want := "Testing failed because of: Test message"

	if got != want {
		t.Errorf("Want: '%s' but got: '%s'", want, got)
	}
}
