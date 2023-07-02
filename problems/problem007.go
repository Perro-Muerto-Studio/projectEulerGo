package problems

// XPrime returns the Xth prime number using the
// Sieve of Eratosthenes algorithm

func XPrime(x int) int {
	primes := sieveOfEratosthenes(200000)
	return primes[x-1]
}

func sieveOfEratosthenes(n int) []int {
	var primes []int
	sieve := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if sieve[i] == false {
			primes = append(primes, i)
			for j := i; j <= n; j += i {
				sieve[j] = true
			}
		}
	}
	return primes
}
