package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	expected := "fl"
	prefix := LongestCommonPrefix(strs)
	if prefix != expected {
		t.Fatalf("the longest common prefix should be %s, but is %s", expected, prefix)
	}

	strs = []string{"cir", "car"}
	expected = "c"
	prefix = LongestCommonPrefix(strs)
	if prefix != expected {
		t.Fatalf("the longest common prefix should be %s, but is %s", expected, prefix)
	}
}
