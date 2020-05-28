package args

import (
	"testing"

	e "github.com/hamburghammer/gstat/errors"
)

func TestNewValidationError(t *testing.T) {
	args := Arguments{}
	message := "Error from testing!"

	got := newValidationError(args, message)
	want := ValidationError{e.BaseError{Operation: "Validation", Message: message}, args}

	if got.Error() != want.Error() {
		t.Errorf("Want: %s Got: %s", want.Error(), got.Error())
	}
}
