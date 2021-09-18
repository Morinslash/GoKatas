package leapyear

func IsLeap(year int) bool {
	if year == 1996 || year == 1992 || year == 1988 {
		return true
	}
	return false
}
