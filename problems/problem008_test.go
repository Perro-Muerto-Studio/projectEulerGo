package problems

import "testing"

func TestLargestProductInSeries(t *testing.T) {
	t.Run("test four adjacent digits", func(t *testing.T) {
		// Arrange
		want := uint64(5832)

		// Act
		got := LargestProductInSeries(4)

		// Assert
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

}
