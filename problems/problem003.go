package problems

import (
	"projectEuler/commons"
)

func LargestPrimeFactor(number int64) int64 {
	largestFactor := int64(2)

	if commons.EsPrimo(number) {
		return number
	}

	for divisor := int64(3); divisor*divisor <= number; divisor += 2 {
		if number%divisor == 0 {
			largestFactor = divisor
			number /= divisor
		}
	}

	if number > largestFactor {
		largestFactor = number
	}

	return largestFactor
}
