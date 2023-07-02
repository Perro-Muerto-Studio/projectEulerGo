package problems

import "projectEuler/commons"

// XPrime returns the Xth prime number using the
// Sieve of Eratosthenes algorithm

func XPrime(x int) int {
	primes := commons.CribaDeEratostenes(200000)
	return primes[x-1]
}
