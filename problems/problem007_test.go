package problems

import (
	"testing"
)

func TestXPrime(t *testing.T) {
	// Arrange
	want := 104743

	// Act
	got := XPrime(10_001)

	// Assert
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
