package exponentiation

func NaiveExponentiation(r, p int) (result int) {
	result = 1
	for i := 1; i <= p; i++ {
		result *= r
	}
	return result
}

func BinaryExponentiation(r, p int) int {
	if p == 0 {
		return 1
	} else if p%2 == 1 {
		return r * BinaryExponentiation(r, p-1)
	} else {
		d := BinaryExponentiation(r, p/2)
		return d * d
	}
}

func MinimumExponentiation(r, p int) (result int) {
	result = 1
	for {
		if p%2 == 1 {
			result *= r
		}
		p /= 2
		if p == 0 {
			break
		}
		r *= r
	}
	return result
}
