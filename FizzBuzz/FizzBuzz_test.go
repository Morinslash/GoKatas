package FizzBuzz_test

import (
	"Katas/FizzBuzz"
	"testing"
)

func TestConvertNormalNumbers(t *testing.T) {
	t.Run("should return 1 when 1 provided", func(t *testing.T) {
		got := FizzBuzz.Convert(1)
		expect := "1"
		if got != expect {
			t.Errorf("Got: %q want: %q", got, expect)
		}
	})
	t.Run("should return 2 when 2 provided", func(t *testing.T) {
		got := FizzBuzz.Convert(2)
		expect := "2"
		if got != expect {
			t.Errorf("Got: %q want: %q", got, expect)
		}
	})
}
