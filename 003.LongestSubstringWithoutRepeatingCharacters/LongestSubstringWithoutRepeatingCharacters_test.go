package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "pwwkew"
	expected := 3
	max := LengthOfLongestSubstring(s)
	if max != expected {
		t.Fatalf("should be %d, but is %d", expected, max)
	}

	s = "pw wkew"
	expected = 4
	max = LengthOfLongestSubstring(s)
	if max != expected {
		t.Fatalf("should be %d, but is %d", expected, max)
	}
}
