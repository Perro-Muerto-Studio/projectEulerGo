package problems

func EvenFibonacciNumbers(limit int64) (sum int64) {
	a := int64(1)
	b := int64(1)

	for b < limit {
		if b%2 == 0 {
			sum += b
		}
		h := a + b
		a, b = b, h
	}

	return
}
