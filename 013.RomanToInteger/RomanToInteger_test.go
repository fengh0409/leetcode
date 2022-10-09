package leetcode

import "testing"

func TestRomanToInteger(t *testing.T) {
	s := "IX"
	expected := 9
	sum := RomanToInteger(s)
	if sum != expected {
		t.Fatalf("should be %d,but is %d", expected, sum)
	}

	s = "MCMXCIV"
	expected = 1994
	sum = RomanToInteger(s)
	if sum != expected {
		t.Fatalf("should be %d,but is %d", expected, sum)
	}
}
