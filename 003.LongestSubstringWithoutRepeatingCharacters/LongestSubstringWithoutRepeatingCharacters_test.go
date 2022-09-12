package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	var s = "pwwkew"

	max := LengthOfLongestSubstring(s)
	if max != 3 {
		t.Fatalf("should be 3, but is %d", max)
	}

	s = "pw wkew"

	max = LengthOfLongestSubstring(s)
	if max != 4 {
		t.Fatalf("should be 4, but is %d", max)
	}
}
