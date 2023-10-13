package commons

func EsPrimo(number int64) (primo bool) {
	switch {
	case number <= 1:
	case number == 2:
	case number%2 == 0:
		return false
	default:
		return esPrimo(number)
	}
	return
}

func esPrimo(number int64) bool {
	for i := int64(3); i*i <= number; i += 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}
