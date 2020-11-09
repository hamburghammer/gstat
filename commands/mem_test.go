package commands_test

import (
	"errors"
	"testing"

	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/commands"
	"github.com/shirou/gopsutil/mem"
)

func TestMemExec(t *testing.T) {
	t.Run("should test for flag in Arguments", func(t *testing.T) {

		got, err := commands.Mem{}.Exec(args.Arguments{Mem: false})

		assertNoError(err, t)

		if len(got) != 0 {
			t.Errorf("No result was expected but got: '%s'", string(got))
		}

	})

	t.Run("should return error if reading disk stats creates one", func(t *testing.T) {
		errorStr := "Test error"
		disk := commands.Mem{ReadVirtualMemoryStat: func() (*mem.VirtualMemoryStat, error) {
			return nil, errors.New(errorStr)
		}}

		_, err := disk.Exec(args.Arguments{Mem: true})
		want := errorStr

		assertEqualString(err.Error(), want, t)

	})

	t.Run("should convert to MB in JSON formatt", func(t *testing.T) {
		disk := commands.Mem{ReadVirtualMemoryStat: func() (*mem.VirtualMemoryStat, error) {
			return &mem.VirtualMemoryStat{Total: 100000000, Used: 50000000}, nil
		}}

		got, err := disk.Exec(args.Arguments{Mem: true})
		want := "{\"mem\":{\"used\":47,\"total\":95}}"

		assertNoError(err, t)

		assertEqualString(string(got), want, t)
	})
}
