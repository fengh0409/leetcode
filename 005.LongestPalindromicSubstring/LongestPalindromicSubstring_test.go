package leetcode

import "testing"

func TestLongestPalindromicSubstring(t *testing.T) {
	s := "abccbd"
	result := LongestPalindromicSubstring(s)
	expected := "bccb"
	if result != expected {
		t.Fatalf("should be %v, but is %v", expected, result)
	}
}
