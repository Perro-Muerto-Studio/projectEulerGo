package problems

import (
	"fmt"
	"math"
)

func LargestPrimeFactor(number int64) (largestFactor int) {
	for factor := 2; factor <= int(math.Sqrt(float64(number))); factor++ {
		if number%int64(factor) == 0 {
			largestFactor = factor
			fmt.Println(factor)
		}
	}
	return
}
