package problems

import "testing"

/*
 * Multiples of 3 or 5
 *
 * If we list all the natural numbers below 10 that are multiples of 3
 * or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 *
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */

func TestProblem001(t *testing.T) {
	// Arrange
	given := 10
	want := 23

	// Act
	got := MultiplesOf3or5(10)

	// Asset
	if got != want {
		t.Errorf("got %d, want %d, max %d", got, want, given)
	}

}
