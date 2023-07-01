package problems

import (
	"fmt"
	"testing"
)

func assertTrue6(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumSquareDifference(t *testing.T) {
	squareDifferencesTest := []struct {
		max  int
		want float64
	}{
		{10, 2_640},
		{100, 25_164_150},
	}

	for _, tt := range squareDifferencesTest {
		t.Run(fmt.Sprintf("Testing for the first %d natural numbers", tt.max), func(t *testing.T) {
			assertTrue6(t, SumSquareDifference(tt.max), tt.want)
		})
	}
}
