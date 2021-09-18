package leapyear

func IsLeap(year int) bool {

	if year%100 == 0 && year%400 != 0 {
		return false
	}

	if year%4 == 0 {
		return true
	}

	return false
}
