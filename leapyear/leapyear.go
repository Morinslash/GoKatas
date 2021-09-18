package leapyear

func IsLeap(year int) bool {
	if isDivisible(400, year) || (isDivisible(4, year) && !isDivisible(100, year)) {
		return true
	}

	return false
}

func isDivisible(divisor int, year int) bool {
	return year%divisor == 0
}
