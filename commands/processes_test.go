package commands_test

import (
	"errors"
	"testing"

	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/commands"
	"github.com/stretchr/testify/assert"
)

func TestProcessExec(t *testing.T) {
	t.Run("should test for the Argument for Process", func(t *testing.T) {
		arguments := args.Arguments{Processes: false}
		got, err := commands.Processes{}.Exec(arguments)

		assert.Nil(t, err, "no error expected")
		assert.Equal(t, 0, len(got))
	})

	t.Run("should exit if the getting the process data returns an error", func(t *testing.T) {
		wantErr := errors.New("getting data error")
		arguments := args.Arguments{Processes: true}
		mockProcessData := func() ([]*commands.Process, error) { return nil, wantErr }
		_, err := commands.Processes{ReadProcesses: mockProcessData}.Exec(arguments)

		assert.NotNil(t, err, "an error was expected")
		assert.Equal(t, wantErr, err)
	})

	t.Run("should pass errors from getting cpu infos to output", func(t *testing.T) {
		wantErr := errors.New("getting data error")
		arguments := args.Arguments{Processes: true}
		nameFunc := func() (string, error) { return "", nil }
		cpuErrorProcessFunc := func() (float64, error) { return 0, wantErr }
		process := commands.Process{Pid: 1, Name: nameFunc, CPUPercent: cpuErrorProcessFunc}
		mockProcessData := func() ([]*commands.Process, error) { return []*commands.Process{&process}, nil }

		_, gotErr := commands.Processes{ReadProcesses: mockProcessData}.Exec(arguments)

		assert.Equal(t, wantErr, gotErr)
	})

	t.Run("should sort per cpu value", func(t *testing.T) {
		arguments := args.Arguments{Processes: true}
		nameFunc := func() (string, error) { return "", nil }
		cpuProcessFunc1 := func() (float64, error) { return 0, nil }
		cpuProcessFunc2 := func() (float64, error) { return 1, nil }
		process1 := commands.Process{Pid: 1, Name: nameFunc, CPUPercent: cpuProcessFunc1}
		process2 := commands.Process{Pid: 2, Name: nameFunc, CPUPercent: cpuProcessFunc2}
		mockProcessData := func() ([]*commands.Process, error) { return []*commands.Process{&process1, &process2}, nil }

		got, err := commands.Processes{ReadProcesses: mockProcessData}.Exec(arguments)
		want := "{\"Processes\":[{\"Name\":\"\",\"Pid\":2,\"CPU\":1},{\"Name\":\"\",\"Pid\":1,\"CPU\":0}]}"

		assert.Nil(t, err)

		assert.Equal(t, want, string(got))
	})

	t.Run("should return max 10 entries", func(t *testing.T) {
		arguments := args.Arguments{Processes: true}
		nameFunc := func() (string, error) { return "", nil }
		cpuProcessFunc := func() (float64, error) { return 0, nil }

		processes := make([]*commands.Process, 12)
		for i := 0; i <= 11; i++ {
			processes[i] = &commands.Process{Pid: int32(i), Name: nameFunc, CPUPercent: cpuProcessFunc}
		}

		mockProcessData := func() ([]*commands.Process, error) { return processes, nil }

		got, err := commands.Processes{ReadProcesses: mockProcessData}.Exec(arguments)
		want := "{\"Processes\":[{\"Name\":\"\",\"Pid\":0,\"CPU\":0},{\"Name\":\"\",\"Pid\":1,\"CPU\":0},{\"Name\":\"\",\"Pid\":2,\"CPU\":0},{\"Name\":\"\",\"Pid\":3,\"CPU\":0},{\"Name\":\"\",\"Pid\":4,\"CPU\":0},{\"Name\":\"\",\"Pid\":5,\"CPU\":0},{\"Name\":\"\",\"Pid\":6,\"CPU\":0},{\"Name\":\"\",\"Pid\":7,\"CPU\":0},{\"Name\":\"\",\"Pid\":8,\"CPU\":0},{\"Name\":\"\",\"Pid\":9,\"CPU\":0}]}"

		assert.Nil(t, err)

		assert.Equal(t, want, string(got))
	})
}
