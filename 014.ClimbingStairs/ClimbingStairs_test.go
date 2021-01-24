package leetcode

import "testing"

func TestClimbingStairs(t *testing.T) {
	n := 3
	result := ClimbingStairs(n)
	if result != 3 {
		t.Fatalf("should be 3, but is %d", result)
	}

	n = 5
	result = ClimbingStairs(n)
	if result != 8 {
		t.Fatalf("should be 8, but is %d", result)
	}
}
