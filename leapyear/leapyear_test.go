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

func TestYearIsNotLeapIfNotDivisibleBy4(t *testing.T) {
	var testCases = []testCase{
		{name: "return false if year 1997", year: 1997, expected: false},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := leapyear.IsLeap(test.year)
			assertGotEqualsCaseExpected(t, got, test)
		})
	}
}

func TestYearIsLeapIfDivisibleBy4(t *testing.T) {
	var testCases = []testCase{
		{name: "return true if year 1996", year: 1996, expected: true},
		{name: "return true if year 1992", year: 1992, expected: true},
		{name: "return true if year 1988", year: 1988, expected: true},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := leapyear.IsLeap(test.year)
			assertGotEqualsCaseExpected(t, got, test)
		})
	}
}

func TestYearIsLeapIfDivisibleBy400(t *testing.T) {
	var testCases = []testCase{
		{name: "return true if year 1600", year: 1600, expected: true},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := leapyear.IsLeap(test.year)
			assertGotEqualsCaseExpected(t, got, test)
		})
	}
}

func TestYearIsNotLeapIfDivisibleBy100ButNotBy400(t *testing.T) {
	var testCases = []testCase{
		{name: "return false if year 1800", year: 1800, expected: false},
		{name: "return false if year 1900", year: 1900, expected: false},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := leapyear.IsLeap(test.year)
			assertGotEqualsCaseExpected(t, got, test)
		})
	}
}

func assertGotEqualsCaseExpected(t *testing.T, got bool, test testCase) {
	if got != test.expected {
		t.Errorf("\n Got: %v \n Expected: %v", got, test.expected)
	}
}
