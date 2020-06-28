package commands_test

import (
	"errors"
	"testing"

	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/commands"
	goDisk "github.com/shirou/gopsutil/disk"
)

func TestDiskExec(t *testing.T) {
	t.Run("should test for flag in Arguments", func(t *testing.T) {

		got, err := commands.Disk{}.Exec(args.Arguments{Disk: false})

		assertNoError(err, t)

		if len(got) != 0 {
			t.Errorf("No result was expected but got: '%s'", string(got))
		}

	})

	t.Run("should return error if reading disk stats creates one", func(t *testing.T) {
		errorStr := "Test error"
		disk := commands.Disk{ReadDiskStats: func(s string) (*goDisk.UsageStat, error) {
			return nil, errors.New(errorStr)
		}}

		_, err := disk.Exec(args.Arguments{Disk: true})
		want := errorStr

		assertEqualString(err.Error(), want, t)

	})

	t.Run("should convert to MB in JSON formatt", func(t *testing.T) {
		disk := commands.Disk{ReadDiskStats: func(s string) (*goDisk.UsageStat, error) {
			return &goDisk.UsageStat{Total: 100000000, Used: 50000000}, nil
		}}

		got, err := disk.Exec(args.Arguments{Disk: true})
		want := "{\"Disk\":\"47/95\"}"

		assertNoError(err, t)

		assertEqualString(string(got), want, t)
	})
}

func assertNoError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("No error was expected but got: '%s'", err.Error())
	}
}

func assertEqualString(got, want string, t *testing.T) {
	if got != want {
		t.Errorf("Want: '%s' but got: '%s'", want, got)
	}
}
