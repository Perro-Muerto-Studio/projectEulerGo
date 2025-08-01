package problems

import "testing"

/*
* Even Fibonacci Numbers
*
* Each new term in the Fibonacci sequence is generated by
* adding the previous two terms. By starting with 1 and
* 2, the first 10 terms will be:
*
*     1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
*
* By considering the terms in the Fibonacci sequence
* whose values do not exceed four million, find the sum
* of the even-valued terms.
*
 */

func TestProblem002(t *testing.T) {
	// Arrange
	limit := int64(89)
	want := int64(44)

	// Act
	got := EvenFibonacciNumbers(limit)

	// Assert
	if got != want {
		t.Errorf("got %v, want %v, limit: %v", got, want, limit)
	}

}
