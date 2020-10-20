package commands

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByCPULen(t *testing.T) {
	t.Run("three item inside the array", func(t *testing.T) {
		array := []cpuProcess{{}, {}, {}}
		got := byCPU(array).Len()
		want := 3

		assert.Equal(t, want, got, "they should be equal")
	})

	t.Run("one item inside the array", func(t *testing.T) {
		array := []cpuProcess{{}}
		got := byCPU(array).Len()
		want := 1

		assert.Equal(t, want, got, "they should be equal")
	})

	t.Run("empty array", func(t *testing.T) {
		array := []cpuProcess{}
		got := byCPU(array).Len()
		want := 0

		assert.Equal(t, want, got, "they should be equal")
	})
}

func TestByCPUSwap(t *testing.T) {
	t.Run("swap array items", func(t *testing.T) {
		unswaped := []cpuProcess{{Name: "foo"}, {Name: "bar"}}
		swaped := []cpuProcess{{Name: "bar"}, {Name: "foo"}}

		got := byCPU(unswaped)
		got.Swap(0, 1)
		want := byCPU(swaped)

		assert.Equal(t, want, got)
	})
}

func TestByCPULess(t *testing.T) {
	cpuProcessArray := []cpuProcess{{CPU: 1}, {CPU: 2}}

	t.Run("less on cpu field smaller", func(t *testing.T) {

		got := byCPU(cpuProcessArray).Less(0, 1)
		want := false

		assert.Equal(t, want, got)
	})

	t.Run("less on cpu field bigger", func(t *testing.T) {

		got := byCPU(cpuProcessArray).Less(1, 0)
		want := true

		assert.Equal(t, want, got)
	})
}

func TestGetProcessCPUInfos(t *testing.T) {
	t.Run("transform Process to cpuProcess", func(t *testing.T) {
		nameFunc := func() (string, error) { return "foo", nil }
		cpuProcessFunc := func() (float64, error) { return 0, nil }
		process := Process{Pid: 1, Name: nameFunc, CPUPercent: cpuProcessFunc}

		got, err := getProcessCPUInfos(&process)
		want := &cpuProcess{Name: "foo", Pid: 1, CPU: 0}

		assert.Nil(t, err, "No error expected")
		assert.Equal(t, want, got)
	})

	t.Run("calling name function returns error", func(t *testing.T) {
		wantErr := errors.New("error")

		nameFunc := func() (string, error) { return "", wantErr }
		cpuProcessFunc := func() (float64, error) { return 0, nil }
		process := Process{Pid: 1, Name: nameFunc, CPUPercent: cpuProcessFunc}

		_, gotErr := getProcessCPUInfos(&process)

		assert.NotNil(t, gotErr, "An error was expected")
		assert.Equal(t, wantErr, gotErr)
	})

	t.Run("calling cpuProcess function returns error", func(t *testing.T) {
		wantErr := errors.New("error")

		nameFunc := func() (string, error) { return "", nil }
		cpuProcessFunc := func() (float64, error) { return 0, wantErr }
		process := Process{Pid: 1, Name: nameFunc, CPUPercent: cpuProcessFunc}

		_, gotErr := getProcessCPUInfos(&process)

		assert.NotNil(t, gotErr, "An error was expected")
		assert.Equal(t, wantErr, gotErr)
	})
}

func TestGetProcessesCPUInfos(t *testing.T) {
	err := errors.New("error")

	nameFunc := func() (string, error) { return "foo", nil }
	nameErrFunc := func() (string, error) { return "foo", err }

	cpuProcessFunc := func() (float64, error) { return 0, nil }

	t.Run("transform array of process into an array of cpuProcess", func(t *testing.T) {
		processes := []*Process{{Pid: 2, Name: nameFunc, CPUPercent: cpuProcessFunc}, {Pid: 1, Name: nameFunc, CPUPercent: cpuProcessFunc}}
		got, gotErr := getProcessesCPUInfos(processes)
		want := []cpuProcess{{Name: "foo", CPU: 0, Pid: 2}, {Name: "foo", CPU: 0, Pid: 1}}

		assert.Nil(t, gotErr, "No error expected")
		assert.Equal(t, want, got)
	})

	t.Run("return error directly if one happens", func(t *testing.T) {
		processes := []*Process{{Pid: 2, Name: nameFunc, CPUPercent: cpuProcessFunc}, {Pid: 1, Name: nameErrFunc, CPUPercent: cpuProcessFunc}}
		got, gotErr := getProcessesCPUInfos(processes)
		want := []cpuProcess{{Name: "foo", CPU: 0, Pid: 2}}
		wantErr := err

		assert.NotNil(t, gotErr, "an error was expected")

		assert.Equal(t, wantErr, gotErr)
		assert.Equal(t, want, got)
	})
}
