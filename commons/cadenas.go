package commons 

func Reverse(s string) (r string) {
	for _, v := range s {
		r = string(v) + r
	}

	return
}
