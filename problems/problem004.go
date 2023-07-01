package problems

import "strconv"
import "projectEuler/commons"

func LargestPalindromeProduct() (num int64) {
	for i:=int64(100); i<=999; i++ {
		for j:=int64(100); j<=999; j++ {
			n := i * j
			p := strconv.FormatInt(n, 10)
			r := commons.Reverse(p)
			if p == r {
				if n > num {
					num = n
				}
			}
		}
	}

	return
}
