package proc

import (
	"testing"
	"time"
)

func mockGetCPUTotal(interval time.Duration, percpu bool) ([]float64, error) {
	return []float64{0.2}, nil
}

func TestCPUTotal(t *testing.T) {
	totalCPUCannel := make(chan float64)
	go TotalCPU(totalCPUCannel, mockGetCPUTotal)
	got := <-totalCPUCannel
	want := 0.2

	if got != want {
		t.Errorf("want %f but got %f", want, got)
	}
}
