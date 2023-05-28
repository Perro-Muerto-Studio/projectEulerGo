package problems

import "testing"

/*
 * Largest Prime Factor
 *
 * The prime factors of 13195 are 5, 7, 13 and 29.
 *
 * What is the largest prime factor of the number 600_851_475_143 ?
 *
 */

func TestProblem003(t *testing.T) {
	t.Run("test largest prime factor 13_195", func(t *testing.T) {
		// Arrange
		number := int64(13_195)
		want := int64(29)

		// Act
		got := LargestPrimeFactor(number)

		// Assert
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

	})
}
