package problems

import "testing"

/*
 * Special Pythagorean Triplet
 * Problem 9

   A  Pythagorean triplet is a set of three natural numbers, a<b<c, for which,
   a²+b²=c²

   For example, 3²+4²=9-49=25=5².

   There exists exactly one Pythagorean triplet for which a+b+c=1000.

   Find the product a*b*c.
*/

func TestSpecialPythagoreanTriplet(t *testing.T) {
	t.Run("test first pytagorean triplet", func(t *testing.T) {
		// Arrange
		want := uint(12)

		// Act
		got := SpecialPythagoreanTriplet(12)

		// Assert
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
