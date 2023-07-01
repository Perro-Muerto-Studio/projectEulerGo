package problems

import (
	"math"
)

// SumSquareDifference calcula la diferencia entre la suma de los cuadrados y
// el cuadrado de la suma de los primeros números naturales indicados por max.
func SumSquareDifference(max int) float64 {
	var i, sum1, sum2 float64
	for i = 0; int(i) <= max; i++ {
		sum1 = math.Pow(i, 2) + sum1
		sum2 = i + sum2
	}
	return math.Pow(sum2, 2) - sum1
}
