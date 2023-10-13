package problems

func SpecialPythagoreanTriplet(suma uint) uint {
	var sumanMil uint

	for a := uint(1); a < suma; a++ {
		for b := uint(1); b < suma; b++ {
			for c := uint(1); c < suma; c++ {
				sumas := a + b + c
				if (a < b) && (b < c) {
					catetos := a*a + b*b
					hipotenusa := c * c
					if catetos == hipotenusa {
						if sumas == suma {
							sumanMil = a * b * c
						}
					}
				}
			}
		}
	}
	return sumanMil
}
