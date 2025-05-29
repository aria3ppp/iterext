package iterext_test

import (
	"slices"
	"testing"

	"github.com/aria3ppp/iterext"
)

func TestMap(t *testing.T) {
	slc := []int{1, 2, 3, 4, 5}

	ptrs := iterext.Map(slices.Values(slc), func(v int) *int {
		return &v
	})
	ptrsSlice := slices.Collect(ptrs)

	if len(ptrsSlice) != len(slc) {
		t.Errorf("expected slice length %d, got %d", len(slc), len(ptrsSlice))
	}
	for i, v := range slc {
		if *ptrsSlice[i] != v {
			t.Errorf("expected value %d at index %d, got %d", v, i, *ptrsSlice[i])
		}
	}
}
