package leetcode

import "testing"

func TestRemoveDuplicatesFromSortArray(t *testing.T) {
	nums := []int{0, 0, 1, 1, 2}
	n := RemoveDuplicatesFromSortArray(nums)
	if n != 3 {
		t.Fatalf("should be 3, but is %d", n)
	}

	nums = []int{0, 1, 2}
	n = RemoveDuplicatesFromSortArray(nums)
	if n != 3 {
		t.Fatalf("should be 3, but is %d", n)
	}
}
