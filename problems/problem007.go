package problems

import "projectEuler/commons"

func XPrime(max int) int {
	i := 0
	count := 0

	for count < max {
		if commons.EsPrimo(int64(i)) {
			count += 1
		}
		i += 1
	}

	return i - 1
}
