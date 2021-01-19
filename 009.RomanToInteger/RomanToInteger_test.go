package leetcode

import "testing"

func TestRomanToInteger(t *testing.T) {
	s := "IX"
	sum := RomanToInteger(s)
	if sum != 9 {
		t.Fatalf("%s should be 9", s)
	}

	s = "MCMXCIV"
	sum = RomanToInteger(s)
	if sum != 1994 {
		t.Fatalf("%s should be 1994", s)
	}
}
