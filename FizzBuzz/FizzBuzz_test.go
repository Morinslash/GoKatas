package FizzBuzz_test

import (
	"Katas/FizzBuzz"
	"testing"
)

func TestConvertNormalNumbersShould(t *testing.T) {
	var testsValues = []struct {
		name   string
		input  int
		expect string
	}{
		{"return 1 when 1 provided", 1, "1"},
		{"return 2 when 2 provided", 2, "2"},
		{"return 4 when 4 provided", 4, "4"},
	}
	for _, testValue := range testsValues {
		t.Run(testValue.name, func(t *testing.T) {
			got := FizzBuzz.Convert(testValue.input)
			if got != testValue.expect {
				t.Errorf("Got: %q want: %q", got, testValue.expect)
			}
		})
	}
}
