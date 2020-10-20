package commands_test

import (
	"errors"
	"testing"

	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/commands"
)

func TestNewResults(t *testing.T) {
	args := args.Arguments{}
	got := commands.NewResult(args)
	want := commands.Result{Arguments: args}

	if !got.ResultEquals(want) {
		t.Errorf("Want: '%v' but got: '%v'", want, got)
	}
}

func TestResultEquals(t *testing.T) {
	t.Run("should be equals", func(t *testing.T) {
		r1 := commands.Result{}
		r2 := commands.Result{}

		got := r1.ResultEquals(r2)
		want := true

		if got != want {
			t.Error("True was expected but got false")
		}
	})

	t.Run("should not be equals if args differ", func(t *testing.T) {
		r1 := commands.Result{Arguments: args.Arguments{CPU: true}}
		r2 := commands.Result{}

		got := r1.ResultEquals(r2)
		want := false

		if got != want {
			t.Error("True was expected but got false")
		}
	})
}

func TestExecCommands(t *testing.T) {
	t.Run("should return string array without opening and closing brackets", func(t *testing.T) {
		arguments := args.Arguments{}
		result := commands.NewResult(arguments)

		mE := mockExecutor{}
		mE.mockExec = func(args.Arguments) ([]byte, error) { return []byte("{test}"), nil }

		executors := []commands.Executor{mE}

		got := result.ExecCommands(executors)
		want := "test"

		if got.Collection.Results[0] != want {
			t.Errorf("Want: '%s' but got: '%s'", want, got.Collection.Results[0])
		}
	})

	t.Run("should return error array", func(t *testing.T) {
		arguments := args.Arguments{}
		result := commands.NewResult(arguments)

		mE := mockExecutor{}
		mE.mockExec = func(args.Arguments) ([]byte, error) { return make([]byte, 0), errors.New("test error") }

		executors := []commands.Executor{mE}

		got := result.ExecCommands(executors)
		want := "test error"

		if got.Collection.Errs[0].Error() != want {
			t.Errorf("Want: '%s' but got: '%s'", want, got.Collection.Results[0])
		}
	})

	t.Run("should return empty result if it gets an empty byte array", func(t *testing.T) {
		arguments := args.Arguments{}
		result := commands.NewResult(arguments)

		mE := mockExecutor{}
		mE.mockExec = func(args.Arguments) ([]byte, error) { return make([]byte, 0), nil }

		executors := []commands.Executor{mE}

		got := len(result.ExecCommands(executors).Collection.Results)
		want := 0

		if got != want {
			t.Errorf("Want: '%d' but got: '%d' entries", want, got)
		}
	})
}

type mockExecutor struct {
	mockExec func(args.Arguments) ([]byte, error)
}

func (mE mockExecutor) Exec(args args.Arguments) ([]byte, error) {
	return mE.mockExec(args)
}
