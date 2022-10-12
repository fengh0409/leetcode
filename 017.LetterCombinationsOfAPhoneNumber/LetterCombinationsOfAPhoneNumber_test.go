package leetcode

import "testing"

func TestLetterCombinationsOfAPhoneNumber(t *testing.T) {
	digits := "23"
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	result := LetterCombinationsOfAPhoneNumber(digits)
	for i := 0; i < len(expected); i++ {
		if expected[i] != result[i] {
			t.Fatalf("the result should be %v, but is %v", expected, result)
		}
	}

	digits = ""
	expected = []string{}
	result = LetterCombinationsOfAPhoneNumber(digits)
	if len(result) != 0 {
		t.Fatalf("the result length should be %v, but is %v", expected, result)
	}
}
