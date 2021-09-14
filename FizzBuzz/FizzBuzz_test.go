package FizzBuzz_test

import (
	"Katas/FizzBuzz"
	"testing"
)

func TestConvertNormalNumbersShould(t *testing.T) {
	var testsValues = []struct {
		testName string
		input    int
		expected string
	}{
		{"return 1 when 1 provided", 1, "1"},
		{"return 2 when 2 provided", 2, "2"},
		{"return 4 when 4 provided", 4, "4"},
	}
	for _, testValue := range testsValues {
		t.Run(testValue.testName, func(t *testing.T) {
			got := FizzBuzz.Convert(testValue.input)
			if got != testValue.expected {
				t.Errorf("Got: %q want: %q", got, testValue.expected)
			}
		})
	}
}

func TestConvertNumberDivisibleBy3Should(t *testing.T) {
	var testsValues = []struct {
		testName string
		input    int
		expected string
	}{
		{"return Fizz when 3 provided", 3, "Fizz"},
		{"return Fizz when 6 provided", 6, "Fizz"},
		{"return Fizz when 9 provided", 9, "Fizz"},
	}
	for _, testValue := range testsValues {
		t.Run(testValue.testName, func(t *testing.T) {
			got := FizzBuzz.Convert(testValue.input)
			if got != testValue.expected {
				t.Errorf("Got: %q want: %q", got, testValue.expected)
			}
		})
	}
}

func TestConvertNumberDivisibleBy5Should(t *testing.T) {
	var testsValues = []struct {
		testName string
		input    int
		expected string
	}{
		{"return Buzz when 5 provided", 5, "Buzz"},
		{"return Buzz when 10 provided", 10, "Buzz"},
		{"return Buzz when 20 provided", 20, "Buzz"},
	}
	for _, testValue := range testsValues {
		t.Run(testValue.testName, func(t *testing.T) {
			got := FizzBuzz.Convert(testValue.input)
			if got != testValue.expected {
				t.Errorf("Got: %q want: %q", got, testValue.expected)
			}
		})
	}
}

func TestConvertNumberDivisibleByBoth3and5Should(t *testing.T) {
	var testsValues = []struct {
		testName string
		input    int
		expected string
	}{
		{"return FizzBuzz when 15 provided", 15, "FizzBuzz"},
		{"return FizzBuzz when 30 provided", 30, "FizzBuzz"},
		{"return FizzBuzz when 45 provided", 45, "FizzBuzz"},
	}
	for _, testValue := range testsValues {
		t.Run(testValue.testName, func(t *testing.T) {
			got := FizzBuzz.Convert(testValue.input)
			if got != testValue.expected {
				t.Errorf("Got: %q want: %q", got, testValue.expected)
			}
		})
	}
}
