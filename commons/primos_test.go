package commons

import (
	"strconv"
	"testing"
)

func assertTrue(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPrimos(t *testing.T) {

	primosTest := []struct {
		number int
		primo  bool
	}{
		{61, true},
		{120121, true},
		{120157, true},
		{120181, true},
		{120193, true},
		{120199, true},
		{120209, true},
		{120223, true},
		{120233, true},
		{120247, true},
		{120277, true},
	}

	for _, tt := range primosTest {
		t.Run(strconv.Itoa(tt.number), func(t *testing.T) {
			assertTrue(t, esPrimo(int64(tt.number)), tt.primo)
		})
	}
}
