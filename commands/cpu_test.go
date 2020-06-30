package commands_test

import (
	"errors"
	"testing"
	"time"

	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/commands"
)

func TestExec(t *testing.T) {
	t.Run("should not run if no args are given", func(t *testing.T) {
		args := args.Arguments{CPU: false}

		got, err := commands.CPU{}.Exec(args)

		assertNoError(err, t)
		if len(got) != 0 {
			t.Error("Got something even though it was not expected")
		}

	})

	t.Run("should return one float", func(t *testing.T) {
		args := args.Arguments{CPU: true}

		readCPUStat := func(interval time.Duration, percpu bool) ([]float64, error) {
			return []float64{float64(0)}, nil
		}

		got, err := commands.CPU{ReadCPUStat: readCPUStat}.Exec(args)
		want := "{\"CPU\":0}"

		assertNoError(err, t)

		if string(got) != want {
			t.Errorf("Want '%s' but got '%s'", want, string(got))
		}
	})

	t.Run("should return wrapped error", func(t *testing.T) {
		args := args.Arguments{CPU: true}

		readCPUStat := func(interval time.Duration, percpu bool) ([]float64, error) {
			return []float64{}, errors.New("Testing error")
		}

		_, got := commands.CPU{ReadCPUStat: readCPUStat}.Exec(args)
		want := "CPUReading failed because of: Testing error"

		assertEqualString(got.Error(), want, t)
	})
}
