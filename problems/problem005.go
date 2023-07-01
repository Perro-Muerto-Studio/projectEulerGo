package problems

func SmallestMultiple(max int) (number int64) {
	number = 1
	for k := 1; k < max; k++ {
		if number % int64(k) > 0 {
			for j := 1; j < max; j++ {
				if (number * int64(j)) % int64(k) == 0 {
					number = number * int64(j)
					break
				}
			}
		}
	}

	return
}
