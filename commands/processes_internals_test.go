package commands

import (
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

// func TestByCPUSwap(t *testing.T) {
// 	t.Run("epmty array", func(t *testing.T) {
// 		unswaped := []cpuProcess{{Name: "foo"}, {Name: "bar"}}
// 		swaped := []cpuProcess{{Name: "bar"}, {Name: "foo"}}
// 		got := byCPU(unswaped)
// 		got.Swap(0, 1)
// 		want := byCPU(swaped)

// 		if got != want {
// 			t.Error("does not equal")
// 		}
// 	})
// }
