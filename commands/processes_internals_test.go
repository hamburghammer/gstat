package commands

import "testing"

func TestByCPULen(t *testing.T) {
	t.Run("three item inside the array", func(t *testing.T) {
		array := []cpuProcess{{}, {}, {}}
		got := byCPU(array).Len()
		want := 3

		assertEqualInt(got, want, t)
	})

	t.Run("one item inside the array", func(t *testing.T) {
		array := []cpuProcess{{}}
		got := byCPU(array).Len()
		want := 1

		assertEqualInt(got, want, t)
	})

	t.Run("epmty array", func(t *testing.T) {
		array := []cpuProcess{}
		got := byCPU(array).Len()
		want := 0

		assertEqualInt(got, want, t)
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

func assertEqualInt(got, want int, t *testing.T) {
	if got != want {
		t.Errorf("Want: '%d' but got: '%d'", want, got)
	}
}
