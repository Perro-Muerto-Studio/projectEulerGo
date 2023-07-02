package commons

func CribaDeEratostenes(n int) []int {
	var primos []int
	criba := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if criba[i] == false {
			primos = append(primos, i)
			for j := i; j <= n; j += i {
				criba[j] = true
			}
		}
	}
	return primos
}
