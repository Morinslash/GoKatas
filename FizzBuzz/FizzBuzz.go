package FizzBuzz

import "strconv"

func Convert(number int) string {
	if divisibleBy(number, 15) {
		return "FizzBuzz"
	}
	if divisibleBy(number, 3) {
		return "Fizz"
	}
	if divisibleBy(number, 5) {
		return "Buzz"
	}
	return strconv.Itoa(number)
}

func divisibleBy(number int, factor int) bool {
	return number%factor == 0
}
