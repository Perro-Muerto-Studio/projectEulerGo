package problems

import (
	"projectEuler/commons"
)

func SummationOfPrimes(limit int64) int64 {
	sum := int64(2)
	for i := int64(3); i <= limit; i += 2 {
		if commons.EsPrimo(i) {
			sum += int64(i)
		}
	}
	return sum
}
