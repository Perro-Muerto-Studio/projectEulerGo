package problems

import (
	"fmt"
	"testing"
)

func assertTrue(t testing.TB, max int, got, want int64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSmallestMultiple(t *testing.T) {

	smallestMultipleTest := []struct {
		max  int
		want int64
	}{
		{10, 2_520},
		{20, 232_792_560},
	}

	for _, tt := range smallestMultipleTest {
		t.Run(fmt.Sprintf("Buscando max:=%v", tt.max), func(t *testing.T) {
			assertTrue(t, tt.max, SmallestMultiple(tt.max), tt.want)
		})
	}

}
