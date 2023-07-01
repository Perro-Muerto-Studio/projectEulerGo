package problems

import "testing"

func TestProblem004(t *testing.T) {
	t.Run("test Largest Palindrome Product", func(t *testing.T) {
		// Arrange
		want := int64(906609)

		// Act
		got := LargestPalindromeProduct()

		// Assert
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
