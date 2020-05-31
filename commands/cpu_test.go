package commands_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/commands"
)

func TestCPUTotal(t *testing.T) {

	t.Run("should return one float", func(t *testing.T) {
		orig := os.Getenv("HOST_PROC")
		os.Setenv("HOST_PROC", "./testdata/proc")

		got, err := commands.CPU{}.TotalCPU()
		want := 0.000000

		if err != nil {
			t.Errorf("Non error was expected but following occurred: %w", err)
		}

		if got != want {
			t.Errorf("Want '%f' but got '%f'", want, got)
		}
		os.Setenv("HOST_PROC", orig)
	})

	// Deactivated due to parallel running test clashing with the env setup
	// t.Run("should catch custom error", func(t *testing.T) {
	// 	orig := os.Getenv("HOST_PROC")
	// 	os.Setenv("HOST_PROC", "./testdata/empty")

	// 	_, got := commands.CPU{}.TotalCPU()
	// 	want := "CPUReading failed because of: No CPU data was found. Please check the HOST_PROC env to point to the right directory."

	// 	if got == nil {
	// 		t.Errorf("An error was expected but not nil")
	// 	}
	// 	if got.Error() != want {
	// 		t.Errorf("Want '%s' but got '%s'", want, got.Error())
	// 	}
	// 	os.Setenv("HOST_PROC", orig)
	// })
}

func TestExec(t *testing.T) {
	orig := os.Getenv("HOST_PROC")
	os.Setenv("HOST_PROC", "./testdata/proc")

	got, err := commands.CPU{}.Exec(args.Arguments{CPU: true})
	want := []byte("{\"CPU\":0}")

	if err != nil {
		t.Errorf("There was an unexpected error: %s", err)
	}

	if !bytes.Equal(got[:], want[:]) {
		t.Errorf("want: '%s' but got: '%s'", string(want), string(got))
	}
	os.Setenv("HOST_PROC", orig)
}
