package problems

import (
	"testing"
)

/*
 * Problem 010
   Summation of Primes

   The sum of the primes below 10 is 2+3+5+7=17.

   Find the sum of all the primes below two million.
 *
*/

func TestSummationOfPrimes(t *testing.T) {
	t.Run("The sum of primes below 10", func(t *testing.T) {
		// Arrange
		want := int64(17)

		// Act
		got := SummationOfPrimes(10)

		// Assert
		if want != got {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
