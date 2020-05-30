package proc

import (
	"os"
	"testing"

	"github.com/shirou/gopsutil/cpu"
)

func TestCPUTotal(t *testing.T) {
	orig := os.Getenv("HOST_PROC")
	os.Setenv("HOST_PROC", "testdata/proc")

	totalCPUCannel := make(chan float64)
	go TotalCPU(totalCPUCannel, cpu.Percent)
	got := <-totalCPUCannel
	want := 0.000000

	if got != want {
		t.Errorf("want %f but got %f", want, got)
	}
	os.Setenv("HOST_PROC", orig)
}
