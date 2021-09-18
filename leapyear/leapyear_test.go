package leapyear_test

import (
	"Katas/leapyear"
	"testing"
)

type testCase struct {
	name     string
	year     int
	expected bool
}

func TestLeapYearShould(t *testing.T) {
	var testCases = []testCase{
		{name: "return false if year 1997", year: 1997, expected: false},
		{name: "return true if year 1996", year: 1996, expected: true},
		{name: "return true if year 1992", year: 1992, expected: true},
		{name: "return true if year 1988", year: 1988, expected: true},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := leapyear.IsLeap(test.year)
			if got != test.expected {
				t.Errorf("Got: %v \n Expected: %v", got, test.expected)
			}
		})
	}
}
