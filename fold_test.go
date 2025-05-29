package iterext_test

import (
	"slices"
	"strconv"
	"testing"

	"github.com/aria3ppp/iterext"
)

func TestFold(t *testing.T) {
	slc := []int{1, 2, 3, 4, 5}
	mp := make(map[string]int, len(slc))

	mp = iterext.Fold(slices.Values(slc), mp, func(m map[string]int, v int) map[string]int {
		m[strconv.Itoa(v)] = v
		return m
	})

	if len(mp) != len(slc) {
		t.Errorf("expected map length %d, got %d", len(slc), len(mp))
	}
	for _, v := range slc {
		if _, ok := mp[strconv.Itoa(v)]; !ok {
			t.Errorf("expected key %s not found in map", strconv.Itoa(v))
		}
	}
}
