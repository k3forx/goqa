package slice_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/goqa/slice"
)

func TestRun(t *testing.T) {
	expectedS1 := [4]int{0, 2, 3, 3}
	expectedS2 := []int{0, 2, 3, 3, 3}

	s1, s2 := slice.Run()

	if diff := cmp.Diff(expectedS1, s1); diff != "" {
		t.Errorf("s1 mismatch (-want, +got):\n%s", diff)
	}
	if diff := cmp.Diff(expectedS2, s2); diff != "" {
		t.Errorf("s2 mismatch (-want, +got):\n%s", diff)
	}
}
