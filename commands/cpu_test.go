package commands_test

import (
	"os"
	"testing"

	"github.com/hamburghammer/gstat/commands"
)

func TestCPUTotal(t *testing.T) {

	t.Run("should return one float", func(t *testing.T) {
		orig := os.Getenv("HOST_PROC")
		os.Setenv("HOST_PROC", "./testdata/proc")

		got, err := commands.TotalCPU()
		want := 0.000000

		if err != nil {
			t.Errorf("Non error was expected but following occurred: %w", err)
		}

		if got != want {
			t.Errorf("Want '%f' but got '%f'", want, got)
		}
		os.Setenv("HOST_PROC", orig)
	})

	t.Run("should catch custom error", func(t *testing.T) {
		orig := os.Getenv("HOST_PROC")
		os.Setenv("HOST_PROC", "./testdata/empty")

		_, got := commands.TotalCPU()
		want := "CPUReading failed because of: No CPU data was found. Please check the HOST_PROC env to point to the right directory."

		if got == nil {
			t.Errorf("An error was expected but not nil")
		}
		if got.Error() != want {
			t.Errorf("Want '%s' but got '%s'", want, got.Error())
		}
		os.Setenv("HOST_PROC", orig)
	})

}
