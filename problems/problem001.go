package problems

func MultiplesOf3or5(max int) (sum int) {
	for v := 0; v < max; v++ {
		if v%3 == 0 || v%5 == 0 {
			sum += v
		}
	}
	return
}
