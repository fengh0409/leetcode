package leetcode

import "testing"

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	expected := 5
	result := FindKthLargest(nums, k)
	if result != expected {
		t.Fatalf("should be %d,but is %d", expected, result)
	}

	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	expected = 4
	result = FindKthLargest(nums, k)
	if result != expected {
		t.Fatalf("should be %d,but is %d", expected, result)
	}
}
